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
export const recExpandButton = writable<boolean>(false);
export const zeroRecs = writable<boolean>(true);

// Azure Resources variables
export const existingResources = writable<string[]>([]);
export const resourcesGeneratedState = writable<boolean>(false);

// Azure VM states
export const existingVMs = writable<string[]>([]);
export const zeroVMs = writable<boolean>(true);
export const vmNames = writable<string[]>([]);
export const resGroups = writable<string[]>([]);
export const operatingSystems = writable<string[]>([]);
export const adminUsernames = writable<string[]>([]);
export const networkInterfaces = writable<string[]>([]);
export const osDisks = writable<string[]>([]);
export const dataDisks = writable<string[]>([]);

// Threat Model states
export const threatModelButtonState = writable<boolean>(false);
export const threatModelGeneratedState = writable<boolean>(false);
export const threatModelGeneratedActual = writable<boolean>(false);

// Threat Model ThreatProcess positions
export const currentposxmod = writable<number>(0);
export const currentposx = writable<number>(200);
export const currentposy = writable<number>(100);