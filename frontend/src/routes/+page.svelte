<script>
    import { onMount } from "svelte";
    import  Message  from "./Message.svelte"

    let socket;
    let data;
    onMount(() => {
        socket = new WebSocket("ws://localhost:8080/ws/chat")
        socket.addEventListener("open", ()=> {
            console.log("Opened");
        })
        socket.addEventListener("message", (event)=> {
             data = JSON.parse(event.data);
        })
    });


</script>

<h1>Un espion xml basique.</h1>

<Message {data} />
