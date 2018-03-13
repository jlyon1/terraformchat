<template>
  <div>
  <div class="section bg">
    <div class="title light">Chatroom: {{title}}</div>
  </div>
  <div class="section">
    <div class="container">
      <div class="chat">
        sdafsaf
      </div>
      <div class="level txt">
        <input class="input"></input>
        <button class="button mg">Send</button>
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
      title: "Unknown"
    }
  }
}
</script>

<style scoped>
.light{
  font-weight: 400;
}
.bg{
  background-color: #74b9ff44;
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
