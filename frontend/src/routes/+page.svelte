<script>
    import { onMount } from "svelte";
    import  Message  from "./Message.svelte"
    import ResetButton from "./ResetButton.svelte";

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

<Message {data} {message_data} />
<ResetButton />
