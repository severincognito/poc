<script>
    import Button, { Label } from '@smui/button';

    let count = 0;

    function resetCounter() {
        const response = fetch('http://localhost:8080/resetCounter');
        let test = response.json();
    }
</script>

<Button on:click={resetCounter} variant="raised">
    <Label>Reset Counter</Label>
</Button>