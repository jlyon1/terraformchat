<template>
  <div>
  <div class="section bg">
    <div class="title light">Chatroom: {{title}}</div>
  </div>
  <div class="section">
    <div class="container">
      <div class="chat">
        <p v-for="str in chat.Log" v-text="str"></p>
      </div>
      <div class="level txt">
        <input v-model="txt" class="input"></input>
        <button class="button mg" @click="send">Send</button>
      </div>
    </div>
  </div>
</div>
</template>

<script>


export default {
  name: 'Main',
  data () {
    return {
      title: "Unknown",
      txt : "",
      chat: ""
    }
  },
  methods: {
    send: function(){
      var val = {"text":this.txt}
      this.txt = ""
      fetch('/send', {
        method: 'post',
        body: JSON.stringify(val)
      }).then(function(response) {
        return response.text();
      }).then(function(data) {
        console.log(data);
      });
    },
    updateChat: function(){
      let el = this;
      fetch("/chat").then(function(data){
        return data.json()
      }).then(function(resp){
        el.chat = resp;
      })
    }
  },
  mounted: function() {
    let el = this;
    setInterval(el.updateChat, 1000);
    fetch("/name").then(function(data){
      return data.text()
    }).then(function(resp){
      el.title = resp;
    })
  }
}
</script>

<style scoped>
.light{
  font-weight: 400;
}
.bg{
  background-color: #74b9ff11;
  border-style: solid;
  border-width: 1px;
  border-color: #eee;
}
.txt{
  margin-top: 15px;

}
.mg{
  margin-left: 5px;
}
.chat{
  border-style: solid;
  border-width: 1px;
  border-radius: 4px;
  border-color: #eee;
  background-color: #dfe6e922;
  padding: 15px;
  overflow-y: scroll;
  height: 500px;
}
</style>
