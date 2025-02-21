import { writable } from 'svelte/store';

export const loggedInUser = writable('');
export const selectedSubscription = writable('');
export const subidfromGin = writable('');

export const existingSubscriptions = writable<string[]>([]);

export const subscriptionButtonState = writable<boolean>(false);
export const subscriptionIsSelectedState = writable<boolean>(false);
export const threatModelGeneratedState = writable<boolean>(false);
