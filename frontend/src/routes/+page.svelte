<script>
    import { onMount } from "svelte";
    import  Message  from "./Message.svelte"
    import  MessageOut  from "./MessageOut.svelte"
    import ResetButton from "./ResetButton.svelte";
    import LayoutGrid, { Cell } from "@smui/layout-grid";
    import Fab from '@smui/fab';
    import { Label } from '@smui/common';

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

<LayoutGrid >
    <Cell order=1 span={2}>
        <div class="demo-cell"><Fab extended><Label>Un simulateur XML basique.</Label></Fab></div>
    </Cell>
    <Cell order=0 span={2}>
        <Cell order=1>
            <div class="demo-cell"><ResetButton /></div>
        </Cell>
        <Cell order=2>
            <div class="demo-cell"><Message {data} {message_data} /></div>
        </Cell>

        <Cell order=2>
            <div class="demo-cell"><MessageOut /></div>
       </Cell>
    </Cell>
</LayoutGrid>


<style>
    .demo-cell {
      /*  height: 60px;*/
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 1em;
    }
</style>