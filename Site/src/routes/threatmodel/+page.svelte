<script lang="ts">
    import { tick } from "svelte";
    import { get } from "svelte/store"
    // Functions from the main svelte page to process threat modeling functions.
    import * as FunctionPersist from "../../stores/persistentfunctions.js";
    // Svelte stores for existing resources.
    import { existingResources, operatingSystems, vmNames, adminUsernames, networkInterfaces, resGroups } from "../../stores/persistentsession.js";
    // Svelte stores for threat model.
    import { threatModelGeneratedState, threatModelGeneratedActual } from "../../stores/persistentsession.js";
    // Svelte stores for selected subscription.
    import { selectedSubscriptionID, selectedSubscriptionName } from "../../stores/persistentsession.js";
    // Svelte stores for current ThreatProcess x and y coordinates.
    import { currentposxmod, currentposx, currentposy } from "../../stores/persistentsession.js";

// Actual components.
class ThreatProcess {
    // Inside drawing text properties.
    readonly font: string = "24px Arial";
    readonly textAlign: CanvasTextAlign = "center";
    readonly textBaseline: CanvasTextBaseline = "middle";
    readonly fillStyle: string = "black";
    readonly radiusP3: number = 40;
    readonly radiusP4: number = 0;
    readonly radiusP5: number = 2 * Math.PI
}

class ThreatInteractor {
    // Shape properties.
    readonly lineWidth: number = 2;
    //readonly
    readonly textBaseline = "top";
}

class ThreatDataStore {
    readonly lineWidth: number = 4;
    readonly textBaseline: CanvasTextBaseline = "middle";
    readonly fillStyle: string = "black";
}

class ThreatDataFlow {

}

class ThreatTrustBoundary {
    readonly setLineDash: Array<number> = [6];
    readonly strokeStyle: string = "red";
    readonly trustBoundaryHeight: number = 1300;
    readonly trustBoundaryWidth: number = 600;
}

const threatProcess = new ThreatProcess;
const threatInteractor = new ThreatInteractor;
const threatDataStore = new ThreatDataStore;
const threatDataFlow = new ThreatDataFlow;
const threatTrustBoundary = new ThreatTrustBoundary;


function generateThreatProcessObject(ctx: CanvasRenderingContext2D) {
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
            ctx.strokeText($vmNames[i], (get(currentposx) + get(currentposxmod)), get(currentposy), 100);
            ctx.strokeStyle = "blue";
            ctx.strokeText("OS: " + $operatingSystems[i], (get(currentposx) + get(currentposxmod)), get(currentposy) - 20, 75);
            ctx.strokeStyle = "black";
            if ($networkInterfaces[i] !== undefined && $networkInterfaces[i] !== null) {
                ctx.strokeStyle = "green";
                const NICstr = ($networkInterfaces[i].split("/")).slice(-1);
                ctx.strokeText("NIC: " + NICstr, (get(currentposx) + get(currentposxmod)), get(currentposy) + 20, 200);
                ctx.strokeStyle = "black";
            }
            if ($adminUsernames[i].includes("adm")) {
                ctx.strokeStyle = "red";
                ctx.strokeText("Weak ADM Username: " + $adminUsernames[i], (get(currentposx) + get(currentposxmod)), get(currentposy) + 40, 200);
                ctx.strokeStyle = "black";
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
    threatModelGeneratedActual.set(true);
    await tick();
    ctx = canvas.getContext("2d") as CanvasRenderingContext2D;
    ctx.setLineDash([6]);
    ctx.strokeStyle = "red";
    ctx.strokeRect(30, 30, 1300, 600);
    ctx.setLineDash([]);
    ctx.strokeText("Threat Boundary - Resource Group: " + $resGroups[0], 20, 20, 500);
    ctx.strokeStyle = "black";
    generateThreatProcessObject(ctx);
    //for (let i = 0; i < $vmNames.length; i++) {
        //generateThreatProcessObject(ctx);
    //}
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