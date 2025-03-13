<script lang="ts">
    // Functions from the main svelte page to process threat modeling functions.
    import { azureResourcesRequest } from "../../stores/persistentfunctions.js";
    // Svelte stores for existing resources.
    import { existingResources } from "../../stores/persistentsession.js";
    // Svelte stores for selected subscription.
    import { selectedSubscriptionID, selectedSubscriptionName } from "../../stores/persistentsession.js";

interface VirtualMachine {
    vmName: string;
    resGroup: string;
    operatingSystem: string;
    adminUsername: string;
    networkInterface: string;
}

// Actual components.
class ThreatProcess {
    // Inside drawing text properties.
    readonly font: string = "16px Arial";
    readonly textAlign: CanvasTextAlign = "center";
    readonly textBaseline: CanvasTextBaseline = "middle";
    readonly fillstyle: string = "black";
    // Radius properties.
    readonly radiusP1: number = 95;
    readonly radiusP2: number = 50;
    readonly radiusP3: number = 40;
    readonly radiusP4: number = 0;
    readonly radiusP5: number = 2 * Math.PI

    draw(ctx: CanvasRenderingContext2D | null): void {
        if (ctx) {
            ctx.beginPath();
            ctx.arc(this.radiusP1, this.radiusP2, this.radiusP3, this.radiusP4, this.radiusP5);
            ctx.stroke();
            ctx.font = this.font;
            ctx.textAlign = this.textAlign;
            ctx.textBaseline = this.textBaseline;
            ctx.fillStyle = this.fillstyle;
        }
    }
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

function DrawThreatProcess(ctx: CanvasRenderingContext2D | null): void {


}

function generateThreatProcessObject(ctx: CanvasRenderingContext2D) {
    const threatProcess = new ThreatProcess;
    if (ctx) {
        ctx.arc(threatProcess.radiusP1, threatProcess.radiusP2, threatProcess.radiusP3, threatProcess.radiusP4, threatProcess.radiusP5);
        ctx.stroke();
        ctx.font = threatProcess.font;
        ctx.textAlign = threatProcess.textAlign;
        ctx.textBaseline = threatProcess.textBaseline;
        ctx.fillStyle = threatProcess.fillstyle;
    }
}

function generateThreatInteractorObject(ctx: CanvasRenderingContext2D) {
    const threatInteractor = new ThreatInteractor;
    if (ctx) {

    }
}

// Generates the canvas elements for the threat model and puts them into an array of stored HTMLCanvasElements, after their context is updated.
export async function generateThreatModel() {
    await azureResourcesRequest();
    let virtualMachineRepository: VirtualMachine[] = [];

    for (let i = 0; i < $existingResources.length; i++) {
        let virtualmachineAttributes: VirtualMachine = {
            vmName: $existingResources[i][i],
            resGroup: $existingResources[i][i+1],
            operatingSystem: $existingResources[i][i+2],
            adminUsername: $existingResources[i][i+3],
            networkInterface: $existingResources[i][i+4]
        }
        virtualMachineRepository.push(virtualmachineAttributes);
    }

    for (let i = 0; i < virtualMachineRepository.length; i++) {
        let ele: HTMLCanvasElement = document.getElementById("vm" + i) as HTMLCanvasElement;
        let ctx = ele.getContext("2d");
        let circle = new ThreatProcess;
        circle.draw(ctx);

        }
    }
</script>