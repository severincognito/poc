<script>
    import Button, { Label } from '@smui/button';
    import Snackbar from '@smui/snackbar';

    let snackbarWithoutClose = Snackbar;
    let result = "";

    async function resetCounter() {
        const response = await fetch('http://localhost:8080/resetCounter');
        let test = await response.json();
        result = test.result;
        snackbarWithoutClose.open();
    }
</script>

<Button on:click={resetCounter} variant="raised">
    <Label>Reset Counter</Label>
</Button>

<Snackbar bind:this={snackbarWithoutClose}>
    <Label>Counter reset result:  {result}.</Label>
</Snackbar>