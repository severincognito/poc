<script>
    export let data = {};
    export let message_data = {};
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';

</script>

<DataTable table$aria-label="People list" style="max-width: 100%;">
    <Head>
        <Row>
            <Cell>Counter</Cell>
            <Cell>Title</Cell>
            <Cell>Message</Cell>
            <Cell numeric>Timestamp</Cell>
        </Row>
    </Head>
    <Body>
    <Row>
        <Cell>{data.Counter}</Cell>
        <Cell>{message_data.Title}</Cell>
        <Cell>{message_data.Content}</Cell>
        <Cell numeric>{message_data.Timestamp}</Cell>
    </Row>
    </Body>
</DataTable>

