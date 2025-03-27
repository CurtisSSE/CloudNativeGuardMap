import { tick } from 'svelte';
import { get } from 'svelte/store';
// Svelte auth related stores.
import { currentposx, currentposxmod, currentposy, loggedInUser, nsgcurrentposx, nsgcurrentposy, subidfromGin } from '../stores/persistentsession.js';
// Svelte stores for threat model related variables.
import { threatModelGeneratedState, threatModelButtonState } from "../stores/persistentsession.js"
// Svelte stores for Virtual Machine states.
import { existingVMs, zeroVMs, vmNames, resGroups, operatingSystems, adminUsernames, networkInterfaces, osDisks, dataDisks } from "../stores/persistentsession.js"
// Svelte stores for Virtual Network states.
import { existingVNs, zeroVNs, vnNames, vnResGroups, vnIPAddresses, vnAddressPrefixes } from "../stores/persistentsession.js"
// Svelte stores for Virtual Network Interface states.
import { existingVNIs, zeroVNIs, vniNames, vniPrivateIPs, vniPublicIPIDs } from '../stores/persistentsession.js'
// Svelte stores for Public IP states.
import { existingPIPs, zeroPIPs, pipPublicIPIDs, pipResGroups, actualPublicIPs } from '../stores/persistentsession.js';
// Svelte stores for Network Security Group states.
import { existingNSGs, zeroNSGs, nsgNames, nsgResGroups, nsgAttachedNIs } from '../stores/persistentsession.js';
// Svelte stores for Resources.
import { resourcesGeneratedState } from "../stores/persistentsession.js";
// Svelte stores for Subscriptions.
import { existingSubscriptions, subscriptionButtonState, selectedSubscriptionName, selectedSubscriptionID, subscriptionIsSelectedState } from "../stores/persistentsession.js";
// Svelte stores for Advisor.
import { existingRecommendations, recExpandButton, zeroRecs, advisorRecommendationsGeneratedState, recName, recID, shortDesc, shortSol, longDesc, actionsDesc, actionsType, actionsCaption, actionsLink, actionsMetaID, impactfromAlert, impactedField, impactedValue, potentialBenefits } from "../stores/persistentsession.js";
// Svelte stores for button states.
import { expandedAdvisorButtonIdx } from "../stores/persistentsession.js"

// Helper variables
const splitVar = " | "

export async function azureLogin() {
    const response = await fetch('http://localhost:5000/auth-login', { method: 'POST' });
    if (response.ok) {
        let result = await response.json();
        let strresult = JSON.stringify(result.user.username);
        loggedInUser.set(strresult.replace(/"/g, ''));
    }
    await tick();
}

export async function azureLogout() {
    const response = await fetch('http://localhost:5000/auth-logout', { method: 'POST' });
    if (response.ok) {
        await response.json();
        loggedInUser.set('');
        selectedSubscriptionName.set('');
        selectedSubscriptionID.set('');
        subidfromGin.set('');
        existingSubscriptions.set([]);
        subscriptionButtonState.set(false);
        subscriptionIsSelectedState.set(false);
        advisorRecommendationsGeneratedState.set(false);
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
        expandedAdvisorButtonIdx.set(null);
        recExpandButton.set(false);
        zeroRecs.set(true);
        resourcesGeneratedState.set(false);
        zeroVMs.set(true);
        existingVMs.set([]);
        vmNames.set([]);
        resGroups.set([]);
        operatingSystems.set([]);
        adminUsernames.set([]);
        networkInterfaces.set([]);
        osDisks.set([]);
        dataDisks.set([]);
        zeroVNs.set(true);
        existingVNIs.set([]);
        vniNames.set([]);
        vniPrivateIPs.set([]);
        vniPublicIPIDs.set([]);
        zeroPIPs.set(true);
        existingPIPs.set([]);
        pipPublicIPIDs.set([]);
        pipResGroups.set([]);
        actualPublicIPs.set([]);
        zeroNSGs.set(true);
        existingNSGs.set([]);
        nsgNames.set([]);
        nsgResGroups.set([]);
        nsgAttachedNIs.set([]);
        threatModelButtonState.set(false);
        threatModelGeneratedState.set(false);
        currentposxmod.set(150);
        currentposx.set(140);
        currentposy.set(150);      
        nsgcurrentposx.set(90);
        nsgcurrentposy.set(90);
    }
    await tick();
}

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
        existingVMs.update(($existingVMs) => [...$existingVMs, ...resourcesFromGin])
        await fetch("http://localhost:5000/vn-request", {method: "POST"}).then(async (response) => {
            if (response.ok) {
                let resourcesvn = await response.json();
                let resourcesFromGinvn = resourcesvn.output;
                existingVNs.update(($existingVNs) => [...$existingVNs, ...resourcesFromGinvn])
                await fetch("http://localhost:5000/vni-request", {method: "POST"}).then(async (response) => {
                    if (response.ok) {
                        let resourcesvni = await response.json();
                        let resourcesFromGinvni = resourcesvni.output;
                        existingVNIs.update(($existingVNIs) => [...$existingVNIs, ...resourcesFromGinvni])
                        await fetch("http://localhost:5000/pip-request", {method: "POST"}).then(async (response) => {
                            if (response.ok) {
                                let resourcespip = await response.json();
                                let resourcesFromGinpip = resourcespip.output;
                                existingPIPs.update(($existingPIPs) => [...$existingPIPs, ...resourcesFromGinpip])
                                await fetch("http://localhost:5000/nsg-request", {method: "POST"}).then(async (response) => {
                                    if (response.ok) {
                                        let resourcesnsg = await response.json();
                                        let resourcesFromGinnsg = resourcesnsg.output;
                                        existingNSGs.update(($existingNSGs) => [...$existingNSGs, ...resourcesFromGinnsg])
                                    }
                                })
                            }
                        })
                    }
                })
            }
        })
        };
    })
    await tick();
}

export async function azureReleaseResources() {
    existingVMs.set([]);
    azureReleaseVirtualMachines();
    azureReleaseVirtualNetworks();
    azureReleaseVirtualNetworkInterfaces();
    azureReleasePublicIPs();
    azureReleaseNSGs();
    currentposx.set(150);
    currentposxmod.set(140);
    currentposy.set(150);
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
    if (get(existingVMs).length > 0) {
        zeroVMs.set(false);
        for (let i = 0; i < get(existingVMs).length; i++) {
            let eVM = get(existingVMs)[i]
            let vmShape = eVM.split(" ");
            vmNames.update(($vmNames) => [...$vmNames, vmShape[0]]);
            resGroups.update(($resGroups) => [...$resGroups, vmShape[1]]);
            operatingSystems.update(($operatingSystems) => [...$operatingSystems, vmShape[2]]);
            adminUsernames.update(($adminUsernames) => [...$adminUsernames, vmShape[3]]);
            networkInterfaces.update(($networkInterfaces) => [...$networkInterfaces, vmShape[4]]);
            osDisks.update(($osDisks) => [...$osDisks, vmShape[5]]);
            dataDisks.update(($dataDisks) => [...$dataDisks, vmShape[6]]);
            }
    } else {
        zeroVMs.set(true);
    }
    await tick(); 
}

export async function azureReleaseVirtualMachines() {
    zeroVMs.set(true);
    vmNames.set([]);
    resGroups.set([]);
    operatingSystems.set([]);
    adminUsernames.set([]);
    networkInterfaces.set([]);
}

export async function azureSetResourceVirtualNetworks() {
    await azureReleaseVirtualNetworks();
    if (get(existingVNs).length > 0) {
        zeroVNs.set(false);
        for (let i = 0; i < get(existingVNs).length; i++) {
            let eVN = get(existingVNs)[i]
            let vnShape = eVN.split(" ");
            vnNames.update(($vnNames) => [...$vnNames, vnShape[0]]);
            vnResGroups.update(($vnResGroups) => [...$vnResGroups, vnShape[1]]);
            vnIPAddresses.update(($vnIPAddresses) => [...$vnIPAddresses, vnShape[2]]);
            vnAddressPrefixes.update(($vnAddressPrefixes) => [...$vnAddressPrefixes, vnShape[3]]);
        }
    } else {
        zeroVNs.set(true);
    }
    await tick();
}


export async function azureReleaseVirtualNetworks() {
    zeroVNs.set(true);
    vnNames.set([]);
    vnResGroups.set([]);
    vnIPAddresses.set([]);
    vnAddressPrefixes.set([]);
}

export async function azureSetResourceVirtualNetworkInterfaces() {
    await azureReleaseVirtualNetworkInterfaces();
    if (get(existingVNIs).length > 0) {
        zeroVNIs.set(false);
        for (let i = 0; i < get(existingVNIs).length; i++) {
            let eVNI = get(existingVNIs)[i]
            let vniShape = eVNI.split(" ");
            vniNames.update(($vniNames) => [...$vniNames, vniShape[0]]);
            vniPrivateIPs.update(($vniPrivateIPs) => [...$vniPrivateIPs, vniShape[1]]);
            vniPublicIPIDs.update(($vniPublicIPIDs) => [...$vniPublicIPIDs, vniShape[2]]);
        }
    } else {
        zeroVNIs.set(true);
    }
    await tick();
}

export async function azureReleaseVirtualNetworkInterfaces() {
    zeroVNIs.set(true);
    vniNames.set([]);
    vniPrivateIPs.set([]);
    vniPublicIPIDs.set([]);
}

export async function azureSetResourcePublicIPs() {
    await azureReleasePublicIPs();
    if (get(existingPIPs).length > 0) {
        zeroPIPs.set(false);
        for (let i = 0; i < get(existingPIPs).length; i++) {
            let ePIP = get(existingPIPs)[i]
            let pipShape = ePIP.split(" ");
            pipPublicIPIDs.update(($pipPublicIPIDs) => [...$pipPublicIPIDs, pipShape[0]]);
            pipResGroups.update(($pipResGroups) => [...$pipResGroups, pipShape[1]]);
            actualPublicIPs.update(($actualPublicIPs) => [...$actualPublicIPs, pipShape[2]]);
        }
    } else {
        zeroPIPs.set(true);
    }
}

export async function azureReleasePublicIPs() {
    zeroPIPs.set(true);
    pipPublicIPIDs.set([]);
    pipResGroups.set([]);
    actualPublicIPs.set([]);
}


export async function azureSetResourceNSGs() {
    await azureReleaseNSGs();
    if (get(existingNSGs).length > 0) {
        zeroNSGs.set(false);
        for (let i = 0; i < get(existingNSGs).length; i++) {
            let eNSG = get(existingNSGs)[i]
            let nsgShape = eNSG.split(" ");
            nsgNames.update(($nsgNames) => [...$nsgNames, nsgShape[0]]);
            nsgResGroups.update(($nsgResGroups) => [...$nsgResGroups, nsgShape[1]]);
            nsgAttachedNIs.update(($nsgAttachedNIs) => [...$nsgAttachedNIs, nsgShape[2]]);
        }
    } else {
        zeroNSGs.set(true);
    }
}

export async function azureReleaseNSGs() {
    zeroNSGs.set(true);
    nsgNames.set([]);
    nsgResGroups.set([]);
    nsgAttachedNIs.set([]);
}

export async function azureSetAdvisorRecommendations() {
    await azureReleaseAdvisorRecommendations();
    await azureAdvisorRequest();
    if (get(existingRecommendations).length > 0) {
        zeroRecs.set(false);
        for (let i = 0; i < get(existingRecommendations).length; i++) {
            const advisorsections = get(existingRecommendations)[i].split(splitVar);
            if (advisorsections[10] == 'Security') {
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
            }}
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