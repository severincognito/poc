<script>
    import { onMount } from "svelte";
    import  Message2  from "./Message2.svelte"
    let test = [];

    let socket;
    let title, message, timestamp;
    onMount(() => {
        socket = new WebSocket("ws://localhost:8080/ws/chat")
        socket.addEventListener("open", ()=> {
            console.log("Opened");
        })
        socket.addEventListener("message", (event)=> {
             let data = JSON.parse(event.data);
             title = data.Title;
             message = data.Content;
             timestamp = data.Timestamp;
        })
    });


</script>

<h1>Un espion xml basique.</h1>

<Message2 title={title} message={message} timestamp={timestamp}/>
