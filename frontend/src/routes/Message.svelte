<script>
    export let data = {};
    export let message_data = {};
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';

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
            <Cell>Counter</Cell>
            <Cell>{data.Counter}</Cell>
        </Row>
        <Row>
            <Cell>Title</Cell>
            <Cell>{message_data.Title}</Cell>
        </Row>
        <Row>
            <Cell>Content</Cell>
            <Cell>{message_data.Content}</Cell>
        </Row>
        <Row>
            <Cell>Timestamp</Cell>
            <Cell numeric>{message_data.Timestamp}</Cell>
        </Row>
    </Body>
</DataTable>
