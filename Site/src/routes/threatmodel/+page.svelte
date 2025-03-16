<script lang="ts">
    import { tick } from "svelte";
    // Functions from the main svelte page to process threat modeling functions.
    import * as FunctionPersist from "../../stores/persistentfunctions.js";
    // Svelte stores for existing resources.
    import { existingResources, operatingSystems, vmNames, adminUsernames, networkInterfaces, resGroups } from "../../stores/persistentsession.js";
    // Svelte stores for threat model.
    import { threatModelGeneratedState, threatModelGeneratedActual } from "../../stores/persistentsession.js";
    // Svelte stores for selected subscription.
    import { selectedSubscriptionID, selectedSubscriptionName } from "../../stores/persistentsession.js";

// Actual components.
class ThreatProcess {
    // Inside drawing text properties.
    readonly font: string = "16px Arial";
    readonly textAlign: CanvasTextAlign = "center";
    readonly textBaseline: CanvasTextBaseline = "middle";
    readonly fillstyle: string = "black";
    // Radius properties.
    radiusx: number = 95;
    radiusy: number = 50;
    readonly radiusP3: number = 40;
    readonly radiusP4: number = 0;
    readonly radiusP5: number = 2 * Math.PI
}

class ThreatInteractor {
    // Shape properties.
    readonly lineWidth: number = 10;
    
    //readonly

}

interface ThreatDataStore {

}

interface ThreatDataFlow {

}

interface ThreatTrustBoundary {

}

function generateThreatProcessObject(ctx: CanvasRenderingContext2D) {
    const threatProcess = new ThreatProcess;
    if (ctx) {
        let radiusxmod: number = 0;
        for (let i = 0; i < $vmNames.length; i++) {
            if (i != 0) {
                radiusxmod += 250;
            }
            ctx.save();
            ctx.beginPath();
            ctx.arc((threatProcess.radiusx + radiusxmod), threatProcess.radiusy, threatProcess.radiusP3, threatProcess.radiusP4, threatProcess.radiusP5);
            ctx.stroke();
            ctx.font = threatProcess.font;
            ctx.textAlign = threatProcess.textAlign;
            ctx.textBaseline = threatProcess.textBaseline;
            ctx.fillStyle = threatProcess.fillstyle;
            ctx.fillText($vmNames[i], (threatProcess.radiusx + radiusxmod), threatProcess.radiusy);
            ctx.closePath();
        }
    }
}

function generateThreatInteractorObject(ctx: CanvasRenderingContext2D | null) {
    const threatInteractor = new ThreatInteractor;
    if (ctx) {

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
    ctx = canvas.getContext('2d') as CanvasRenderingContext2D; 
// Drawing begins here.
    
    for (let i = 0; i < $vmNames.length; i++) {
        generateThreatProcessObject(ctx);
        }
    }
    
</script>

{#if $threatModelGeneratedState == true && $threatModelGeneratedActual == false}
<div>
<button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {generateThreatModel();}}>Generate threat model</button>
</div>
{:else if $threatModelGeneratedState == true && $threatModelGeneratedActual == true}

<center>
    <p class="font-bold">Subscription ID: {$selectedSubscriptionID} | Subscription Name: {$selectedSubscriptionName}</p>
    <canvas bind:this={canvas} width="1080" height="720"></canvas></center>
{/if}