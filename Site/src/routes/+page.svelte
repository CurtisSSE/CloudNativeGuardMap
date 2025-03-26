<script lang="ts">
    // Import all persistent state functions.
    import * as FunctionPersist from "../stores/persistentfunctions.js"
    import { tick } from "svelte";
    import { get } from "svelte/store"
    // Svelte stores for Virtual Machine states.
    import { operatingSystems, vmNames, adminUsernames, networkInterfaces, resGroups, osDisks, dataDisks } from "../stores/persistentsession.js";
    // Svelte stores for Virtual Network states.
    import { vnNames, vnResGroups, vnIPAddresses, vnAddressPrefixes } from "../stores/persistentsession.js"
    // Svelte stores for Virtual Network Interface states.
    import { vniNames, vniPrivateIPs, vniPublicIPIDs } from '../stores/persistentsession.js'
    // Svelte stores for Public IP states.
    import { pipPublicIPIDs, pipResGroups, actualPublicIPs } from '../stores/persistentsession.js';
    // Svelte stores for Network Security Group states.
    import { nsgNames, nsgResGroups, nsgAttachedNIs } from '../stores/persistentsession.js';
    // Svelte stores for threat model.
    import { threatModelGeneratedState } from "../stores/persistentsession.js";
    // Svelte stores for current ThreatProcess x and y coordinates.
    import { currentposxmod, currentposx, currentposy } from "../stores/persistentsession.js";
    // Import individual identity states.
    import { loggedInUser } from "../stores/persistentsession.js";
    // Import individual subscription states.
    import { selectedSubscriptionName, selectedSubscriptionID, subscriptionIsSelectedState, subscriptionButtonState, existingSubscriptions } from "../stores/persistentsession.js";
    // Import individual recommendation states.
    import { advisorRecommendationsGeneratedState, recExpandButton, zeroRecs, recName, recID, shortDesc, impactedField, impactfromAlert, impactedValue, expandedAdvisorButtonIdx, existingRecommendations} from "../stores/persistentsession.js";
    // Import individual resources states.
    import { resourcesGeneratedState } from "../stores/persistentsession.js";

    // Actual components.
class ThreatProcess {
    readonly font: string = "24px Arial";
    readonly textAlign: CanvasTextAlign = "center";
    readonly textBaseline: CanvasTextBaseline = "middle";
    readonly fillStyle: string = "black";
    readonly radiusP3: number = 40;
    readonly radiusP4: number = 0;
    readonly radiusP5: number = 2 * Math.PI
}

class ThreatDataStore {
    readonly textBaseline: CanvasTextBaseline = "middle";
    readonly strokeStyle: string = "black";
}

class ThreatTrustBoundary {
    readonly setLineDash: Array<number> = [6];
    readonly strokeStyle: string = "red";

    readonly subscriptionTrustBoundaryHeight: number = 650;
    readonly subscriptionTrustBoundaryWidth: number = 1550;
    readonly subscriptionTrustBoundaryPosx: number = 30;
    readonly subscriptionTrustBoundaryPosy: number = 30;

    readonly resourcegroupTrustBoundaryHeight: number = 550;
    readonly resourcegroupTrustBoundaryWidth: number = 1500;
    readonly resourcegroupTrustBoundaryPosx: number = 60;
    readonly resourcegroupTrustBoundaryPosy: number = 60;

    readonly nsgTrustBoundaryHeight: number = 500;
    readonly nsgTrustBoundaryWidth: number = 400;
    nsgTrustBoundaryPosx: number = 90;
    nsgTrustBoundaryPosy: number = 90;

    readonly threatTrustBoundaryTextMaxWidth: number = 1000;
}
// Helper variables.
const splitVar = " | "
let canvas: HTMLCanvasElement;
let ctx: CanvasRenderingContext2D;
const splitVarNetworkInterface = "/networkInterfaces/"
let threatProcess = new ThreatProcess;
let threatDataStore = new ThreatDataStore;
let threatTrustBoundary = new ThreatTrustBoundary;

function generateVMSubFunction(ctx: CanvasRenderingContext2D) {
    if (ctx) {
        for (let i = 0; i < $vmNames.length; i++) {
            ctx.save();
            if ($networkInterfaces[i].split(splitVarNetworkInterface)[1] === $nsgAttachedNIs[0].split(splitVarNetworkInterface)[1]) {
                ctx.strokeStyle = threatTrustBoundary.strokeStyle;
                ctx.setLineDash(threatTrustBoundary.setLineDash);
                ctx.strokeRect(threatTrustBoundary.nsgTrustBoundaryPosx, threatTrustBoundary.nsgTrustBoundaryPosy, threatTrustBoundary.nsgTrustBoundaryWidth, threatTrustBoundary.nsgTrustBoundaryHeight);
                ctx.setLineDash([]);
                ctx.strokeText("Threat Boundary - Network Security Group: " + $nsgNames[0], threatTrustBoundary?.nsgTrustBoundaryPosx-10, threatTrustBoundary.nsgTrustBoundaryPosy-10, threatTrustBoundary.threatTrustBoundaryTextMaxWidth);
                ctx.strokeStyle = "black";
                threatTrustBoundary.nsgTrustBoundaryPosx = (threatTrustBoundary.nsgTrustBoundaryPosx * 2)
            }
            ctx.beginPath();
            ctx.arc($currentposx + $currentposxmod, $currentposy, threatProcess.radiusP3, threatProcess.radiusP4, threatProcess.radiusP5);
            ctx.stroke();
            ctx.font = threatProcess.font;
            ctx.textAlign = threatProcess.textAlign;
            ctx.textBaseline = threatProcess.textBaseline;
            ctx.fillStyle = threatProcess.fillStyle;
            ctx.strokeText("VM: " + $vmNames[i], $currentposx + $currentposxmod, $currentposy, 200);
            ctx.strokeStyle = "blue";
            ctx.strokeText("OS: " + $operatingSystems[i], $currentposx + $currentposxmod, $currentposy - 50, 200);
            ctx.strokeStyle = "black";
            if ($networkInterfaces[i] !== undefined && $networkInterfaces[i] !== null) {
                ctx.strokeStyle = "green";
                const NICstr = ($networkInterfaces[i].split("/")).slice(-1);
                ctx.strokeText("NIC: " + NICstr, ($currentposx + $currentposxmod), $currentposy + 40, 200);
                ctx.strokeStyle = "black";
            }
            if ($vniPrivateIPs[i] !== undefined && $vniPrivateIPs[i] !== null) {
                ctx.strokeStyle = "black";
                const PrivateIPstr = ($vniPrivateIPs[i])
                ctx.strokeText("Private IP: " + PrivateIPstr, ($currentposx + $currentposxmod), $currentposy + 60, 200);
            }
            if ($actualPublicIPs[i] !== undefined && $actualPublicIPs[i] !== null) {
                ctx.strokeStyle = "black";
                const PublicIPstr = ($actualPublicIPs[i])
                ctx.strokeText("Public IP: " + PublicIPstr, ($currentposx + $currentposxmod), $currentposy + 80, 200);
            }
            if ($adminUsernames[i].includes("adm")) {
                ctx.strokeStyle = "red";
                ctx.strokeText("Weak ADM Username: " + $adminUsernames[i], ($currentposx + $currentposxmod), $currentposy + 60, 150);
                ctx.strokeStyle = "black";
            }
            ctx.closePath();
            ctx.beginPath();
            ctx.strokeStyle = threatDataStore.strokeStyle;
            if ($osDisks[i] !== undefined && $osDisks[i] !== null) {
                ctx.moveTo($currentposx + $currentposxmod, $currentposy + 60)
                ctx.lineTo($currentposx + $currentposxmod, $currentposy + 90);
                ctx.stroke();
                ctx.lineTo($currentposx + 5 + $currentposxmod, $currentposy + 85);
                ctx.stroke();
                ctx.moveTo($currentposx + $currentposxmod, $currentposy + 90);
                ctx.lineTo($currentposx - 5 + $currentposxmod, $currentposy + 85);
                ctx.stroke();
                ctx.moveTo($currentposx + $currentposxmod - 100, $currentposy + 100);
                ctx.lineTo($currentposx + $currentposxmod + 100, $currentposy + 100);
                ctx.stroke();
                ctx.textBaseline = threatDataStore.textBaseline;
                ctx.strokeText("OS Disk: " + $osDisks[i], $currentposx + $currentposxmod, $currentposy + 140, 225);
                ctx.moveTo($currentposx + $currentposxmod - 100, $currentposy + 180);
                ctx.lineTo($currentposx + $currentposxmod + 100, $currentposy + 180);
                ctx.stroke();
            }
            ctx.closePath();
            ctx.beginPath();
            ctx.strokeStyle = threatDataStore.strokeStyle;
            ctx.moveTo($currentposx + $currentposxmod - 100, $currentposy + 240);
            if ($dataDisks[i] !== undefined && $dataDisks[i] !== null && $dataDisks[i] != 'No-Data-Disks') {
                ctx.moveTo($currentposx + $currentposxmod, $currentposy + 200)
                ctx.lineTo($currentposx + $currentposxmod, $currentposy + 230);
                ctx.stroke();
                ctx.lineTo($currentposx + 5 + $currentposxmod, $currentposy + 225);
                ctx.stroke();
                ctx.moveTo($currentposx + $currentposxmod, $currentposy + 230);
                ctx.lineTo($currentposx - 5 + $currentposxmod, $currentposy + 225);
                ctx.stroke();
                ctx.moveTo($currentposx + $currentposxmod - 100, $currentposy + 240);
                ctx.lineTo($currentposx + $currentposxmod + 100, $currentposy + 240);
                ctx.stroke();
                ctx.textBaseline = threatDataStore.textBaseline;
                ctx.strokeText("Attached Disk: " + $dataDisks[i], $currentposx + $currentposxmod, $currentposy + 280, 225);
                ctx.moveTo($currentposx + $currentposxmod - 100, $currentposy + 320);
                ctx.lineTo($currentposx + $currentposxmod + 100, $currentposy + 320);
                ctx.stroke();
                }
                currentposxmod.set($currentposxmod + $currentposxmod);
                currentposx.set($currentposx + $currentposxmod);
                currentposy.set($currentposy);
                ctx.closePath();
                ctx.restore();
        }
    }
}

function generateThreatModelSubFunction(ctx: CanvasRenderingContext2D) {
    if (ctx) {
            generateVMSubFunction(ctx);
    }
}

// Generates the canvas elements for the threat model and puts them into an array of stored HTMLCanvasElements, after their context is updated.
async function generateThreatModel() {
    await FunctionPersist.azureResourcesRequest();
    await FunctionPersist.azureSetResourceVirtualMachines();
    await FunctionPersist.azureSetResourceVirtualNetworks();
    await FunctionPersist.azureSetResourceVirtualNetworkInterfaces();
    await FunctionPersist.azureSetResourcePublicIPs();
    await FunctionPersist.azureSetResourceNSGs();
    threatModelGeneratedState.set(true);

    await tick();
    ctx = canvas.getContext("2d") as CanvasRenderingContext2D;
    ctx.setLineDash(threatTrustBoundary.setLineDash);
    ctx.strokeStyle = threatTrustBoundary.strokeStyle;
    ctx.strokeRect(threatTrustBoundary.subscriptionTrustBoundaryPosx, threatTrustBoundary.subscriptionTrustBoundaryPosy, threatTrustBoundary.subscriptionTrustBoundaryWidth, threatTrustBoundary.subscriptionTrustBoundaryHeight);
    ctx.setLineDash([]);
    ctx.strokeText("Threat Boundary - Subscription: " + $selectedSubscriptionName + " | " + "Subscription ID: " + $selectedSubscriptionID, threatTrustBoundary.subscriptionTrustBoundaryPosx-10, threatTrustBoundary.subscriptionTrustBoundaryPosy-10, threatTrustBoundary.threatTrustBoundaryTextMaxWidth);
    ctx.setLineDash(threatTrustBoundary.setLineDash);
    ctx.strokeRect(threatTrustBoundary.resourcegroupTrustBoundaryPosx, threatTrustBoundary.resourcegroupTrustBoundaryPosy, threatTrustBoundary.resourcegroupTrustBoundaryWidth, threatTrustBoundary.resourcegroupTrustBoundaryHeight);
    ctx.setLineDash([]);
    ctx.strokeText("Threat Boundary - Resource Group: " + $resGroups[0], threatTrustBoundary.resourcegroupTrustBoundaryPosx-10, threatTrustBoundary.resourcegroupTrustBoundaryPosy-10, threatTrustBoundary.threatTrustBoundaryTextMaxWidth);
    ctx.strokeStyle = "black";
    generateThreatModelSubFunction(ctx);
}

async function releaseThreatModel() {
    threatModelGeneratedState.set(false);
    await FunctionPersist.azureReleaseResources();
    threatProcess = new ThreatProcess;
    threatDataStore = new ThreatDataStore;
    threatTrustBoundary = new ThreatTrustBoundary;
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.setTransform(1, 0, 0, 1, 0, 0);
}

</script>

<div>
    {#snippet selectedSubscriptionSnippet()}
        <p class="font-bold">Current selected subscription:</p>{$selectedSubscriptionName}       
    {/snippet}

    {#snippet availableAzureSubscriptionsSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 justify-center text-lg" onclick={FunctionPersist.azureStartSubscriptionsAuth}>Show available Azure Subscriptions</button>
    {/snippet}

    {#snippet eachAvailableAzureSubscriptionsSnippet()}
        <p class="font-bold">Choose an Azure subscription:</p>
        {#each $existingSubscriptions as subscriptions}
        <br/><button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureSetSubscription(subscriptions.split(splitVar)[1], subscriptions.split(splitVar)[0]); FunctionPersist.azureStartAdvisorAuths($selectedSubscriptionID); FunctionPersist.azureStartResourcesAuth($selectedSubscriptionID); FunctionPersist.azureToggleSubscriptionState()}}>{subscriptions}</button><br/>
        {/each}
    {/snippet}

    {#snippet changeSelectedSubscriptionSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureReleaseSubscription(); FunctionPersist.azureToggleThreatModelState(); FunctionPersist.azureReleaseAdvisorRecommendations(); FunctionPersist.azureToggleSubscriptionState()}}>Change Subscription</button>
    {/snippet}

    {#snippet startAdvisorRecommendationsSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => { FunctionPersist.azureToggleAdvisorState(); FunctionPersist.azureSetAdvisorRecommendations();}}>Check Azure Advisor security recommendations</button>
    {/snippet}

    {#snippet startThreatModelSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => { FunctionPersist.azureToggleThreatModelState(); generateThreatModel();}}>Start Threat Modeling</button>
    {/snippet}


    {#snippet displayAdvisorRecommendationsSnippet()}
        <center><br/>
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureReleaseAdvisorRecommendations(); FunctionPersist.azureToggleAdvisorState() }}>Cancel and return to menu</button><br/><br/>
        <p class="font-semibold">Azure Security Posture Recommendations:</p><br/>
        {#if $recExpandButton == true && $zeroRecs == true}
        <p>No security recommendations available for this subscription.</p>
        {:else if $recExpandButton == true && $zeroRecs == false}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureRecButtonHandler()}}>Close Recommendations</button><br/>
        {#each $recName as _, i}
        <br/>
        <div class="my-3">
            <button class="p-4 rounded-lg border border-gray-300 bg-white shadow-sm cursor-pointer hover: bg-gray-50 transition-colors flex justify-between items-center" onclick={() => FunctionPersist.toggleAlertExpand(i)}>
                <div class="flex items-center space-x-3">
                    <span class="bg-blue-100 text-blue-800 px-2 py-1 rounded-md font-medium">Alert #{i+1}</span>
                    <span class="text-gray-700 font-medium">{$shortDesc[i]} | Severity: </span>
                </div>
                <span class="px-2 py-1 rounded-md text-sm font-medium"
                    class:bg-red-100={$impactedField[i] === 'High'}
                    class:text-red-800={$impactedField[i] === 'High'}
                    class:bg-yellow-100={$impactedField[i] === 'Medium'}
                    class:text-yellow-800={$impactedField[i] === 'Medium'}
                    class:bg-green-100={$impactedField[i] === 'Low'}
                    class:text-green-800={$impactedField[i] === 'Low'}>
                    {$impactedField[i]}
                </span>
            </button>

            {#if $expandedAdvisorButtonIdx === i}
            <div class="pt-5 px-16 animate-expandVertical">
                <div class="grid gap-3">
                    <div class="flex">
                        <p class="font-bold w-40">Alert ID:</p>
                        <span class="text-gray-700 font-mono">{$recName[i]}</span>
                    </div>
                    <div class="flex">
                        <p class="font-bold w-40">Resource:</p>
                        <span class="text-gray-700">{$recID[i]}</span>
                    </div>
                    <div class="flex">
                        <p class="font-bold w-40">Risk Type:</p>
                        <span class="text-gray-700">{$impactfromAlert[i]}</span>
                    </div>
                    <div class="flex">
                        <p class="font-bold w-40">Resource Type:</p>
                        <span class="text-gray-700">{$impactedValue[i]}</span>
                    </div>
                </div>
                <br/><br/>
            </div>
            {/if}
        </div>
        {/each}
        {:else if $recExpandButton == false}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureSetAdvisorRecommendations(); FunctionPersist.azureRecButtonHandler()}}>Expand Recommendations</button>
        {/if}
        </center>
    {/snippet}

    {#snippet generateThreatModelSnippet()}
    <center>
        <center><br/>
            <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => { FunctionPersist.azureToggleThreatModelState(); releaseThreatModel(); }}>Cancel and return to menu</button></center>
        <br/>
        <canvas bind:this={canvas} class="border-2 border-black" width="1680" height="800"></canvas>
    </center>
    {/snippet}

    <!-- If the user is logged in and a subscription is not selected. -->
    {#if $loggedInUser != '' && $subscriptionButtonState == false}
    <center>
    {@render availableAzureSubscriptionsSnippet()}
    </center>
    <!-- If the user is logged in, they clicked choose a subscription and a subscription is not yet selected. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == false}
    <center>
    {@render eachAvailableAzureSubscriptionsSnippet()}
    </center>
    <!-- If the user is logged in and has selected a subscription. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == false && $threatModelGeneratedState == false}
    <center>
    {@render selectedSubscriptionSnippet()}<br/><br/>
    {@render changeSelectedSubscriptionSnippet()}<br/><br/>
    {@render startAdvisorRecommendationsSnippet()}<br/><br/>
    {@render startThreatModelSnippet()}
    </center>

    <!-- If the user is logged in, has selected a subscription and has selected to generate security advisor recommendations. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == true && $threatModelGeneratedState == false}
    {@render displayAdvisorRecommendationsSnippet()}

    <!-- If the user is logged in, has selected a subscription and has selected to generate a threat model. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == false && $threatModelGeneratedState == true}
    {@render generateThreatModelSnippet()}

    <!-- If the user is not logged in. -->
    {:else if $loggedInUser == ''}
    <div class="text-center">Not logged in. Please log in above to continue.</div>
    {/if}
</div>
