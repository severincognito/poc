<script>
    import { onMount } from "svelte";
    import  Message  from "./Message.svelte"
    import ResetButton from "./ResetButton.svelte";
    import LayoutGrid, { Cell } from "@smui/layout-grid";

    let socket;
    let data, message_data;
    onMount(() => {
        socket = new WebSocket("ws://localhost:8080/ws/chat")
        socket.addEventListener("open", ()=> {
            console.log("Opened");
        })
        socket.addEventListener("message", (event)=> {
             data = JSON.parse(event.data);
             message_data = data.Message;
        })
    });
</script>

<h1>Un espion xml basique.</h1>

<LayoutGrid>
        <Cell order=1>
            <div class="demo-cell"><ResetButton /></div>
        </Cell>
        <Cell order=0>
            <div class="demo-cell"><Message {data} {message_data} /></div>
        </Cell>
</LayoutGrid>

<style>
    .demo-cell {
        height: 60px;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>