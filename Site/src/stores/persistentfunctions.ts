import { tick } from 'svelte';
import { get } from 'svelte/store';
// Svelte stores for threat model related variables.
import { threatModelGeneratedState, threatModelButtonState, threatModelGeneratedActual, zeroVMs, vmNames, resGroups, operatingSystems, adminUsernames, networkInterfaces} from "../stores/persistentsession.js";
// Svelte stores for Resources.
import { existingResources, resourcesGeneratedState } from "../stores/persistentsession.js";
// Svelte stores for Subscriptions.
import { existingSubscriptions, subscriptionButtonState, selectedSubscriptionName, selectedSubscriptionID, subscriptionIsSelectedState } from "../stores/persistentsession.js";
// Svelte stores for Advisor.
import { recExpandButton, zeroRecs, existingRecommendations, advisorRecommendationsGeneratedState, recName, recID, shortDesc, shortSol, longDesc, actionsDesc, actionsType, actionsCaption, actionsLink, actionsMetaID, impactfromAlert, impactedField, impactedValue, potentialBenefits } from "../stores/persistentsession.js";
// Svelte stores for button states.
import { expandedAdvisorButtonIdx } from "../stores/persistentsession.js"

// Helper variables
const splitVar = " | "

    // Primary Azure SubscriptionsFactoryClient auth handler and returns list of subscriptions for logged in user's tenant.
export async function azureStartSubscriptionsAuth() {
    await fetch('http://localhost:5000/subscriptions-auth', { method: 'POST' })
    .then(async (response) => {
        if (response.ok) {
        let subs = await response.json();
        let subscriptionsfromGin = subs.output
        existingSubscriptions.update(($existingSubscriptions) => [...$existingSubscriptions, ...subscriptionsfromGin]);
        };
    subscriptionButtonState.set(true);
    })
}

// Authentication handler and primary triggering function for additional auths required for diagram, such as ARMAdvisor Client Factory.
export async function azureStartAdvisorAuths(subscriptionid: string) {
    await fetch("http://localhost:5000/advisor-auth", {method: "POST", headers: { "Content-Type": "application/json"}, 
    body: JSON.stringify({ subscriptionid })
    }).then(async (response) => {
        if (response.ok) {
        let subid = await response.json();
        return subid.output;
    };
    })
    await tick();
}

export async function azureAdvisorRequest() {
    await fetch("http://localhost:5000/advisor-request", {method: "POST"}).then(async (response) => {
        if (response.ok) {
        let recs = await response.json();
        let recommendationsFromGin = recs.output;
        existingRecommendations.update(($existingRecommendations) => [...$existingRecommendations, ...recommendationsFromGin]);
        };
    })
    await tick();
}

export async function azureStartResourcesAuth(subscriptionid: string) {
    await fetch("http://localhost:5000/resources-auth", {method: "POST", headers: { "Content-Type": "application/json"},
    body: JSON.stringify({ subscriptionid })
    }).then(async (response) => {
        if (response.ok) {
            let subid = await response.json();
            return subid.output;
        };
    })
    await tick();
}

export async function azureResourcesRequest() {
    await azureReleaseResources();
    await fetch("http://localhost:5000/resources-request", {method: "POST"}).then(async (response) => {
        if (response.ok) {
        let resources = await response.json();
        let resourcesFromGin = resources.output;
        existingResources.update(($existingResources) => [...$existingResources, ...resourcesFromGin])
        };
    })
    await tick();
}

export async function azureReleaseResources() {
    existingResources.set([]);
}

    // Sets the Azure subscription state amongst Svelte stores and other components.
export async function azureSetSubscription(subscriptionname: string, subscriptionid: string) {
    azureReleaseSubscription();
    selectedSubscriptionName.set(subscriptionname);
    selectedSubscriptionID.set(subscriptionid);
    await tick();
    return selectedSubscriptionName;
}

    // Unsets the Azure subscription state amongst Svelte stores and other components.
export async function azureReleaseSubscription() {
    selectedSubscriptionName.set('');
    selectedSubscriptionID.set('');
    await tick();
}

export async function azureToggleSubscriptionState() {
    if (get(subscriptionIsSelectedState) == true) {
        subscriptionIsSelectedState.set(false);
    } else {
        subscriptionIsSelectedState.set(true);
    }
    await tick();
}

export async function azureToggleThreatModelState() {
    if (get(threatModelGeneratedState) == true) {
        threatModelGeneratedState.set(false);
    } else {
        threatModelGeneratedState.set(true);
    }
    await tick();
}

export async function azureSetResourceVirtualMachines() {
    await azureReleaseVirtualMachines();
    await azureReleaseResources();
    await azureResourcesRequest();
    if (get(existingResources).length > 0) {
        zeroVMs.set(false);
        for (let i = 0; i < get(existingResources).length; i++) {
            let vmShape = get(existingResources)[i].split(" ");
            vmNames.update(($vmNames) => [...$vmNames, vmShape[0]]);
            resGroups.update(($resGroups) => [...$resGroups, vmShape[1]]);
            operatingSystems.update(($operatingSystems) => [...$operatingSystems, vmShape[2]]);
            adminUsernames.update(($adminUsernames) => [...$adminUsernames, vmShape[3]]);
            networkInterfaces.update(($networkInterfaces) => [...$networkInterfaces, vmShape[4]]);
            }
    } else {
        zeroVMs.set(true);
    }
    await tick(); 
}

export async function azureReleaseVirtualMachines() {
    zeroVMs.set(true);
    existingResources.set([]);
    vmNames.set([]);
    resGroups.set([]);
    operatingSystems.set([]);
    adminUsernames.set([]);
    networkInterfaces.set([]);
}

export async function azureSetAdvisorRecommendations() {
    await azureReleaseAdvisorRecommendations();
    await azureAdvisorRequest();
    if (get(existingRecommendations).length > 0) {
        zeroRecs.set(false);
        for (let i = 0; i < get(existingRecommendations).length; i++) {
            const advisorsections = get(existingRecommendations)[i].split(splitVar);
            //if (advisorsections[10] == 'Security') {
                recName.update(($recName) => [...$recName, advisorsections[0]]);
                recID.update(($recID) => [...$recID, advisorsections[1].split('/providers/Microsoft.Advisor/rec')[0]]);
                shortDesc.update(($shortDesc) => [...$shortDesc, advisorsections[2]]);
                shortSol.update(($shortSol) => [...$shortSol, advisorsections[3]]);
                longDesc.update(($longDesc) => [...$longDesc, advisorsections[4]]);
                actionsDesc.update(($actionsDesc) => [...$actionsDesc, advisorsections[5]]);
                actionsType.update(($actionsType) => [...$actionsType, advisorsections[6]]);
                actionsCaption.update(($actionsCaption) => [...$actionsCaption, advisorsections[7]]);
                actionsLink.update(($actionsLink) => [...$actionsLink, advisorsections[8]]);
                actionsMetaID.update(($actionsMetaID) => [...$actionsMetaID, advisorsections[9]]);
                impactfromAlert.update(($impactfromAlert) => [...$impactfromAlert, advisorsections[10]]);
                impactedField.update(($impactedField) => [...$impactedField, advisorsections[11]]);
                impactedValue.update(($impactedValue) => [...$impactedValue, advisorsections[12]]);
                potentialBenefits.update(($potentialBenefits) => [...$potentialBenefits, advisorsections[13]]);
            }
    } else {
        zeroRecs.set(true);
    }
    await tick();
}

export async function azureReleaseAdvisorRecommendations() {
    zeroRecs.set(true);
    existingRecommendations.set([]);
    recName.set([]);
    recID.set([]);
    shortDesc.set([]);
    shortSol.set([]);
    longDesc.set([]);
    actionsDesc.set([]);
    actionsType.set([]);
    actionsCaption.set([]);
    actionsLink.set([]);
    actionsMetaID.set([]);
    impactfromAlert.set([]);
    impactedField.set([]);
    impactedValue.set([]);
    potentialBenefits.set([]);
    await tick();
}

export async function azureToggleAdvisorState() {
    if (get(advisorRecommendationsGeneratedState) == true) {
        advisorRecommendationsGeneratedState.set(false);
    } else {
        advisorRecommendationsGeneratedState.set(true);
    }
    await tick();
}

export async function azureRecButtonHandler() {
    if (get(recExpandButton) == false) {
        recExpandButton.set(true);
    } else {
        recExpandButton.set(false);
    }
    await tick();
}

export async function azureToggleResourcesState() {
    if (get(resourcesGeneratedState) == true) {
        resourcesGeneratedState.set(false);
    } else {
        resourcesGeneratedState.set(true);
    }
    await tick();
}

export function toggleAlertExpand(idx: any) {
    expandedAdvisorButtonIdx.set(expandedAdvisorButtonIdx === idx ? null : idx);
}

export async function azureToggleThreatModelActualState() {
    if (get(threatModelGeneratedActual) == true) {
        threatModelGeneratedActual.set(false);
    } else {
        threatModelGeneratedActual.set(true);
    }
    await tick();
}