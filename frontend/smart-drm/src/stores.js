import { writable } from "svelte/store";

export const isConnected = writable(false);
export const provider = writable(undefined);
export const signer = writable(undefined);
export const error = writable("");
