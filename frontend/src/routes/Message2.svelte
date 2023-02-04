<script>
    export let title;
    export let message;
    export let timestamp;
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
</script>

<DataTable table$aria-label="People list" style="max-width: 100%;">
    <Head>
        <Row>
            <Cell>Title</Cell>
            <Cell>Message</Cell>
            <Cell numeric>Timestamp</Cell>
        </Row>
    </Head>
    <Body>
    <Row>
        <Cell>{title}</Cell>
        <Cell>{message}</Cell>
        <Cell numeric>{timestamp}</Cell>
    </Row>
    </Body>
</DataTable>

