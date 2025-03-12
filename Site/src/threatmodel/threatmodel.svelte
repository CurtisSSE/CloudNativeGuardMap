{#snippet threatModel()}
    <script lang="ts">

        const canvas = document.getElementById('ThreatModel') as HTMLCanvasElement

        const ctx = canvas.getContext('2d') as CanvasRenderingContext2D;

        ctx.clearRect(0, 0, canvas.width, canvas.height);

        ctx.fillStyle = '#3498db'

    </script>
{/snippet}