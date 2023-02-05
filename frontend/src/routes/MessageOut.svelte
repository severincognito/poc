<script>
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
    import Button, { Label } from '@smui/button';
    import Textfield from '@smui/textfield';
    import HelperText from '@smui/textfield/helper-text';
    import Snackbar from '@smui/snackbar';

    let snackbarWithoutClose = Snackbar;

    let message_data = {};

    let Content = "";
    let Title ="";
    let Timestamp = 0;

    let result;

    async function transmit () {
        const res = await fetch('http://localhost:8080/transmit', {
            method: 'POST',
            body: JSON.stringify({
                "title": Title,
                "content": Content,
                "timestamp": Timestamp
            })
        })
        let resultJSON = await res.json()
        result = resultJSON.result;
        snackbarWithoutClose.open();
    }

</script>

<DataTable table$aria-label="Message" style="max-width: 100%;">
    <Head>
        <Row>
            <Cell>Field</Cell>
            <Cell>Value</Cell>
        </Row>
    </Head>
    <Body>
        <Row>
            <Cell>Title</Cell>
            <Cell>
                <Textfield bind:value={Title} label="Title">
                    <HelperText slot="helper">Message Title</HelperText>
                </Textfield>
            </Cell>
        </Row>
        <Row>
            <Cell>Content</Cell>
            <Cell>
                    <Textfield bind:value={Content} label="Content">
                    <HelperText slot="helper">Message Content</HelperText>
                    </Textfield>
            </Cell>
        </Row>
        <Row>
            <Cell>Timestamp</Cell>
            <Cell numeric>
                <Textfield bind:value={Timestamp} label="Timestamp" type="number" />
            </Cell>
        </Row>
    </Body>
</DataTable>


<Button on:click={transmit} variant="raised">
    <Label>Transmit</Label>
</Button>

<Snackbar bind:this={snackbarWithoutClose}>
    <Label>Transmit result:  {result}.</Label>
</Snackbar>