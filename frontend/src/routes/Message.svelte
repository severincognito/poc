<script>
    export let data = {};
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
        <Cell>{data.Message.Title}</Cell>
        <Cell>{data.Message.Content}</Cell>
        <Cell numeric>{data.Message.Timestamp}</Cell>
    </Row>
    </Body>
</DataTable>

