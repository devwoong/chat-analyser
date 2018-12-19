<template>
    <div class="chatlog" ref="chatlog" v-on:scroll="scroll">
        <div class="chat_wrapper" v-for="(chat,idx) in resultChats">
            <!-- <p style="font-size:14px">{{chat._source.author}}</p> -->
            <div  :class="[chat._source.author == '회원님'? 'my_chat_node' : 'chat_node']">
            <template v-if="idx == '0'">
                <p style="font-size:14px;">{{chat._source.author}}</p>
            </template>
            <template v-else>
                <p v-if="resultChats[idx-1]._source.author != chat._source.author" style="font-size:14px;">{{chat._source.author}}</p>
            </template>

            <div>
                <p>{{chat._source.contents}}</p>
            </div>
            <template v-if="idx == resultChats.length-1">
                <p style="font-size:10px;">{{chat._source.date}}</p>
            </template>
            <template v-else>
                <p v-if="resultChats[idx+1]._source.author != chat._source.author" style="font-size:10px;">{{chat._source.date}}</p>
            </template>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    created() {
        this.$http.get(`${this.api_url}/date/chat`, {
            params:{
                size:this.size,
                from:(this.currentPage-1)*this.size,
                toDate : this.toDate,
                fromDate : this.fromDate
                },
                headers:{}, timeout:10000})
            .then((result) => {
                this.resultChats = result.data.hits.hits;
            })
    },
    data() {
        return {
            size : 20,
            currentPage : 1,
            resultChats : [],
            scrollSize : 750,    
        }
    },
    methods : {
        scroll(e) {
            var scrollTop = this.$refs.chatlog.scrollTop
            var scrollHeight = this.$refs.chatlog.scrollHeight
            if( scrollHeight - scrollTop <= this.scrollSize){
                this.currentPage++;
                this.search();
            }
        },
        search() {
            this.$http.get(`${this.api_url}/date/chat`, {
                params:{
                    size:this.size,
                    from:(this.currentPage-1)*this.size,
                    toDate : this.toDate,
                    fromDate : this.fromDate
                    },
                    headers:{}, timeout:10000})
                .then(function(result){
                    for(var i=0; i<result.data.hits.hits.length; i++){
                        this.resultChats.push(result.data.hits.hits[i]);
                    }
                }.bind(this))
        }
    },
    props : ['fromDate', 'toDate', 'searchOrder'],
    watch : {
        searchOrder() {
            this.currentPage = 1;
            this.resultChats = [];
            this.search();
        }
    }
}
</script>

<style>
    .chatlog {
        border: 1px solid black;
        background-color: #9bbdd4;
        margin-top: 10px;
        height : 750px;
        overflow: scroll;
    }
    .chat_wrapper {
        display: flow-root;
        border: none;
    }
    .chat_node {
        max-width: 80%;
        float: left;
        border: none;
        margin-left: 3px;
    }
    .my_chat_node {
        max-width: 80%;
        float: right;
        border: none;
        margin-right: 3px;
    }
    .chat_node p {
        color : #556677;
        text-align: center;
        word-break:break-all;
        margin-left: 4px;
        margin-right: 4px;
    }
    .my_chat_node p {
        color : #556677;
        text-align: center;
        word-break:break-all;
        margin-left: 4px;
        margin-right: 4px;
    }
    .chat_node div {
        border-radius: 2em;
        background-color:#ffffff;
    }
    .my_chat_node div {
         border-radius: 2em;
        background-color:#fef01b;
    }
</style>
