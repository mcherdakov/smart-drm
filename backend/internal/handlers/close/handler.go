package close

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

type drmService interface {
	UserAddrToChannelAddr(userAddr string) (string, error)
	CallContractCloseChannel(ctx context.Context, hash []byte, channel common.Address) (*types.Receipt, error)
}

type proofsRepo interface {
	DeleteProof(ctx context.Context, address string) error
	ProofByAddress(ctx context.Context, address string) (*entity.DBProof, error)
}

type Handler struct {
	drm  drmService
	repo proofsRepo
}

func NewHandler(drm drmService, repo proofsRepo) *Handler {
	return &Handler{
		drm:  drm,
		repo: repo,
	}
}

type Request struct {
	Address string `json:"address"`
}

func (h *Handler) Handle(c *gin.Context) {
	rawBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
		return
	}

	var req Request
	if err := json.Unmarshal(rawBody, &req); err != nil {
		c.Error(err)
		return
	}

	chanAddr, err := h.drm.UserAddrToChannelAddr(req.Address)
	if err != nil {
		c.Error(err)
		return
	}

	proof, err := h.repo.ProofByAddress(c, chanAddr)
	if err != nil {
		c.Error(err)
		return
	}

	if proof == nil {
		c.Error(fmt.Errorf("no proof yet"))
		return
	}

	hash, err := hex.DecodeString(proof.Hash[2:])
	if err != nil {
		c.Error(err)
		return
	}

	_, err = h.drm.CallContractCloseChannel(
		c,
		hash,
		common.HexToAddress(chanAddr),
	)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.repo.DeleteProof(c, chanAddr); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}
