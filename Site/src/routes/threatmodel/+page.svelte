<script lang="ts">
    import { tick } from "svelte";
    import { get } from "svelte/store"
    // Functions from the main svelte page to process threat modeling functions.
    import * as FunctionPersist from "../../stores/persistentfunctions.js";
    // Svelte stores for Virtual Machine states.
    import { operatingSystems, vmNames, adminUsernames, networkInterfaces, resGroups, osDisks, dataDisks } from "../../stores/persistentsession.js";
    // Svelte stores for Virtual Network states.
    import { vnNames, vnResGroups, vnIPAddresses, vnAddressPrefixes } from "../../stores/persistentsession.js"
    // Svelte stores for Virtual Network Interface states.
    import { vniNames, vniPrivateIPs, vniPublicIPIDs } from '../../stores/persistentsession.js'
    // Svelte stores for Public IP states.
    import { pipPublicIPIDs, pipResGroups, actualPublicIPs } from '../../stores/persistentsession.js';
    // Svelte stores for Network Security Group states.
    import { nsgNames, nsgResGroups, nsgAttachedNIs } from '../../stores/persistentsession.js';
    // Svelte stores for threat model.
    import { threatModelGeneratedState, threatModelGeneratedActual } from "../../stores/persistentsession.js";
    // Svelte stores for selected subscription.
    import { selectedSubscriptionID, selectedSubscriptionName } from "../../stores/persistentsession.js";
    // Svelte stores for current ThreatProcess x and y coordinates.
    import { currentposxmod, currentposx, currentposy } from "../../stores/persistentsession.js";
const splitVar = "~"
const splitVarNetworkInterface = "/networkInterfaces/"

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

class ThreatInteractor {
    readonly lineWidth: number = 4;
    readonly textBaseline = "middle";
}

class ThreatDataStore {
    readonly textBaseline: CanvasTextBaseline = "middle";
    readonly strokeStyle: string = "black";
}

class EntityProcess {
    readonly placeholder: string = ""
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

const threatProcess = new ThreatProcess;
const threatInteractor = new ThreatInteractor;
const entityProcess = new EntityProcess;
const threatDataStore = new ThreatDataStore;
const threatTrustBoundary = new ThreatTrustBoundary;


function generateVMSubFunction(ctx: CanvasRenderingContext2D) {
    if (ctx) {
        for (let i = 0; i < $vmNames.length; i++) {
            ctx.save();
            if ($networkInterfaces[i].split(splitVarNetworkInterface)[1] === $nsgAttachedNIs[0].split(splitVarNetworkInterface)[1]) {
                ctx.strokeStyle = threatTrustBoundary.strokeStyle;
                ctx.setLineDash(threatTrustBoundary.setLineDash);
                ctx.strokeRect(threatTrustBoundary.nsgTrustBoundaryPosx, threatTrustBoundary.nsgTrustBoundaryPosy, threatTrustBoundary.nsgTrustBoundaryWidth, threatTrustBoundary.nsgTrustBoundaryHeight);
                ctx.setLineDash([]);
                ctx.strokeText("Threat Boundary - Network Security Group: " + $nsgNames[0], threatTrustBoundary.nsgTrustBoundaryPosx-10, threatTrustBoundary.nsgTrustBoundaryPosy-10, threatTrustBoundary.threatTrustBoundaryTextMaxWidth);
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
                ctx.moveTo(get(currentposx) + get(currentposxmod), get(currentposy) + 90);
                ctx.lineTo(get(currentposx) - 5 + get(currentposxmod), get(currentposy) + 85);
                ctx.stroke();
                ctx.moveTo(get(currentposx) + get(currentposxmod) - 100, get(currentposy) + 100);
                ctx.lineTo(get(currentposx) + get(currentposxmod) + 100, get(currentposy) + 100);
                ctx.stroke();
                ctx.textBaseline = threatDataStore.textBaseline;
                ctx.strokeText("OS Disk: " + $osDisks[i], get(currentposx) + get(currentposxmod), get(currentposy) + 140, 225);
                ctx.moveTo(get(currentposx) + get(currentposxmod) - 100, get(currentposy) + 180);
                ctx.lineTo(get(currentposx) + get(currentposxmod) + 100, get(currentposy) + 180);
                ctx.stroke();
            }
            ctx.closePath();
            ctx.beginPath();
            ctx.strokeStyle = threatDataStore.strokeStyle;
            ctx.moveTo(get(currentposx) + get(currentposxmod) - 100, get(currentposy) + 240);
            if ($dataDisks[i] !== undefined && $dataDisks[i] !== null && $dataDisks[i] != 'No-Data-Disks') {
                ctx.moveTo(get(currentposx) + get(currentposxmod), get(currentposy) + 200)
                ctx.lineTo(get(currentposx) + get(currentposxmod), get(currentposy) + 230);
                ctx.stroke();
                ctx.lineTo(get(currentposx) + 5 + get(currentposxmod), get(currentposy) + 225);
                ctx.stroke();
                ctx.moveTo(get(currentposx) + get(currentposxmod), get(currentposy) + 230);
                ctx.lineTo(get(currentposx) - 5 + get(currentposxmod), get(currentposy) + 225);
                ctx.stroke();
                ctx.moveTo(get(currentposx) + get(currentposxmod) - 100, get(currentposy) + 240);
                ctx.lineTo(get(currentposx) + get(currentposxmod) + 100, get(currentposy) + 240);
                ctx.stroke();
                ctx.textBaseline = threatDataStore.textBaseline;
                ctx.strokeText("Attached Disk: " + $dataDisks[i], get(currentposx) + get(currentposxmod), get(currentposy) + 280, 225);
                ctx.moveTo(get(currentposx) + get(currentposxmod) - 100, get(currentposy) + 320);
                ctx.lineTo(get(currentposx) + get(currentposxmod) + 100, get(currentposy) + 320);
                ctx.stroke();
                }
                ctx.closePath();
                ctx.restore();
                currentposxmod.set($currentposxmod + $currentposxmod);
                currentposx.set($currentposx + $currentposxmod);
                currentposy.set($currentposy);
            }
    }
}

function generateThreatModelSubFunction(ctx: CanvasRenderingContext2D) {
    if (ctx) {
            generateVMSubFunction(ctx);
    }
}

let canvas: HTMLCanvasElement;
let ctx: CanvasRenderingContext2D;

// Generates the canvas elements for the threat model and puts them into an array of stored HTMLCanvasElements, after their context is updated.
async function generateThreatModel() {
    await FunctionPersist.azureToggleThreatModelActualState();
    await FunctionPersist.azureResourcesRequest();
    await FunctionPersist.azureSetResourceVirtualMachines();
    await FunctionPersist.azureSetResourceVirtualNetworks();
    await FunctionPersist.azureSetResourceVirtualNetworkInterfaces();
    await FunctionPersist.azureSetResourcePublicIPs();
    await FunctionPersist.azureSetResourceNSGs();
    threatModelGeneratedActual.set(true);

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
    
</script>

{#if $threatModelGeneratedState == true && $threatModelGeneratedActual == false}
<div>
<button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {generateThreatModel(); }}>Generate threat model</button>
</div>
{:else if $threatModelGeneratedState == true && $threatModelGeneratedActual == true}
<center>
    <br/>
    <canvas bind:this={canvas} class="border-2 border-black" width="1680" height="800"></canvas>
</center>
{/if}