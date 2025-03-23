<script lang="ts">
    import { tick } from "svelte";
    import { get } from "svelte/store"
    // Functions from the main svelte page to process threat modeling functions.
    import * as FunctionPersist from "../../stores/persistentfunctions.js";
    // Svelte stores for existing resources.
    import { operatingSystems, vmNames, adminUsernames, networkInterfaces, resGroups, osDisks, dataDisks} from "../../stores/persistentsession.js";
    // Svelte stores for threat model.
    import { threatModelGeneratedState, threatModelGeneratedActual } from "../../stores/persistentsession.js";
    // Svelte stores for selected subscription.
    import { selectedSubscriptionID, selectedSubscriptionName } from "../../stores/persistentsession.js";
    // Svelte stores for current ThreatProcess x and y coordinates.
    import { currentposxmod, currentposx, currentposy } from "../../stores/persistentsession.js";

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

class ThreatDataFlow {
}

class ThreatTrustBoundary {
    readonly setLineDash: Array<number> = [6];
    readonly strokeStyle: string = "red";
    readonly subscriptionTrustBoundaryHeight: number = 500;
    readonly subscriptionTrustBoundaryWidth: number = 1450;
    readonly resourcegroupTrustBoundaryHeight: number = 400;
    readonly resourcegroupTrustBoundaryWidth: number = 1350;
}

const threatProcess = new ThreatProcess;
const threatInteractor = new ThreatInteractor;
const threatDataStore = new ThreatDataStore;
const threatDataFlow = new ThreatDataFlow;
const threatTrustBoundary = new ThreatTrustBoundary;


function generateThreatModelSubFunction(ctx: CanvasRenderingContext2D) {
    if (ctx) {
        for (let i = 0; i < $vmNames.length; i++) {
            if (i != 0) {
                currentposxmod.set(get(currentposxmod) + 250);
            }
            ctx.save();
            ctx.beginPath();
            ctx.arc((get(currentposx) + get(currentposxmod)), get(currentposy), threatProcess.radiusP3, threatProcess.radiusP4, threatProcess.radiusP5);
            ctx.stroke();
            ctx.font = threatProcess.font;
            ctx.textAlign = threatProcess.textAlign;
            ctx.textBaseline = threatProcess.textBaseline;
            ctx.fillStyle = threatProcess.fillStyle;
            ctx.strokeText("VM: " + $vmNames[i], (get(currentposx) + get(currentposxmod)), get(currentposy), 200);
            ctx.strokeStyle = "blue";
            ctx.strokeText("OS: " + $operatingSystems[i], (get(currentposx) + get(currentposxmod)), get(currentposy) - 20, 150);
            ctx.strokeStyle = "black";
            if ($networkInterfaces[i] !== undefined && $networkInterfaces[i] !== null) {
                ctx.strokeStyle = "green";
                const NICstr = ($networkInterfaces[i].split("/")).slice(-1);
                ctx.strokeText("NIC: " + NICstr, (get(currentposx) + get(currentposxmod)), get(currentposy) + 20, 150);
                ctx.strokeStyle = "black";
            }
            if ($adminUsernames[i].includes("adm")) {
                ctx.strokeStyle = "red";
                ctx.strokeText("Weak ADM Username: " + $adminUsernames[i], (get(currentposx) + get(currentposxmod)), get(currentposy) + 40, 150);
                ctx.strokeStyle = "black";
            }
            ctx.closePath();
            ctx.beginPath();
            ctx.strokeStyle = threatDataStore.strokeStyle;
            if ($osDisks[i] !== undefined && $osDisks[i] !== null) {
                ctx.moveTo(get(currentposx) + get(currentposxmod), get(currentposy) + 60)
                ctx.lineTo(get(currentposx) + get(currentposxmod), get(currentposy) + 90);
                ctx.stroke();
                ctx.lineTo(get(currentposx) + 5 + get(currentposxmod), get(currentposy) + 85);
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
        }
        currentposx.set((get(currentposx) + get(currentposxmod) + get(currentposxmod)));
        currentposy.set(get(currentposy));
    }
}

let canvas: HTMLCanvasElement;
let ctx: CanvasRenderingContext2D;

// Generates the canvas elements for the threat model and puts them into an array of stored HTMLCanvasElements, after their context is updated.
async function generateThreatModel() {
    await FunctionPersist.azureToggleThreatModelActualState();
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
    ctx.strokeRect(30, 30, threatTrustBoundary.subscriptionTrustBoundaryWidth, threatTrustBoundary.subscriptionTrustBoundaryHeight);
    ctx.setLineDash([]);
    ctx.strokeText("Threat Boundary - Subscription: " + $selectedSubscriptionName + " | " + "Subscription ID: " + $selectedSubscriptionID, 20, 20, 1000);
    ctx.setLineDash(threatTrustBoundary.setLineDash);
    ctx.strokeRect(60, 60, threatTrustBoundary.resourcegroupTrustBoundaryWidth, threatTrustBoundary.resourcegroupTrustBoundaryHeight);
    ctx.setLineDash([]);
    ctx.strokeText("Threat Boundary - Resource Group: " + $resGroups[0], 50, 50, 500);
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