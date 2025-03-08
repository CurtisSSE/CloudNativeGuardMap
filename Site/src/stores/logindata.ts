import { writable } from 'svelte/store';

// User login states
export const loggedInUser = writable('');

// Subscription selected main states
export const selectedSubscriptionName = writable('');
export const selectedSubscriptionID = writable('');
export const subidfromGin = writable('');

// Subscription feedback from Gin backend
export const existingSubscriptions = writable<string[]>([]);

// Subscription states
export const subscriptionButtonState = writable<boolean>(false);
export const subscriptionIsSelectedState = writable<boolean>(false);

// Threat Model states
export const threatModelGeneratedState = writable<boolean>(false);

// Azure Advisor recommendation variables
export const advisorRecommendationsGeneratedState = writable<boolean>(false);
export const existingRecommendations = writable<string[]>([]);
export const recName = writable<string[]>([]);
export const recID = writable<string[]>([]);
export const shortDesc = writable<string[]>([]);
export const shortSol = writable<string[]>([]);
export const longDesc = writable<string[]>([]);
export const actionsDesc = writable<string[]>([]);
export const actionsType = writable<string[]>([]);
export const actionsCaption = writable<string[]>([]);
export const actionsLink = writable<string[]>([]);
export const actionsMetaID = writable<string[]>([]);
export const impactfromAlert = writable<string[]>([]);
export const impactedField = writable<string[]>([]);
export const impactedValue = writable<string[]>([]);
export const potentialBenefits = writable<string[]>([]);

// Azure Advisor recommendation states
export const expandedAdvisorButtonIdx = writable<number | null>(null);