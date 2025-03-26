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
export const resourcesGeneratedState = writable<boolean>(false);

// Azure VM states
export const zeroVMs = writable<boolean>(true);
export const existingVMs = writable<string[]>([]);
export const vmNames = writable<string[]>([]);
export const resGroups = writable<string[]>([]);
export const operatingSystems = writable<string[]>([]);
export const adminUsernames = writable<string[]>([]);
export const networkInterfaces = writable<string[]>([]);
export const osDisks = writable<string[]>([]);
export const dataDisks = writable<string[]>([]);

// Azure VN states
export const zeroVNs = writable<boolean>(true);
export const existingVNs = writable<string[]>([]);
export const vnNames = writable<string[]>([]);
export const vnResGroups = writable<string[]>([]);
export const vnIPAddresses = writable<string[]>([]);
export const vnAddressPrefixes = writable<string[]>([]);

// Azure VNI states
export const zeroVNIs = writable<boolean>(true);
export const existingVNIs = writable<string[]>([]);
export const vniNames = writable<string[]>([]);
export const vniPrivateIPs = writable<string[]>([]);
export const vniPublicIPIDs = writable<string[]>([]);

// Azure PIP states
export const zeroPIPs = writable<boolean>(true);
export const existingPIPs = writable<string[]>([]);
export const pipPublicIPIDs = writable<string[]>([]);
export const pipResGroups = writable<string[]>([]);
export const actualPublicIPs = writable<string[]>([]);

// Azure NSG states
export const zeroNSGs = writable<boolean>(true);
export const existingNSGs = writable<string[]>([]);
export const nsgNames = writable<string[]>([]);
export const nsgResGroups = writable<string[]>([]);
export const nsgAttachedNIs = writable<string[]>([]);

// Threat Model states
export const threatModelButtonState = writable<boolean>(false);
export const threatModelGeneratedState = writable<boolean>(false);

// Threat Model ThreatProcess positions
export const currentposxmod = writable<number>(150);
export const currentposx = writable<number>(140);
export const currentposy = writable<number>(150);