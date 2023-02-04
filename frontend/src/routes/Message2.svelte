<script>
    export let data = {};
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
        <Cell>{data.Title}</Cell>
        <Cell>{data.Content}</Cell>
        <Cell numeric>{data.Timestamp}</Cell>
    </Row>
    </Body>
</DataTable>

