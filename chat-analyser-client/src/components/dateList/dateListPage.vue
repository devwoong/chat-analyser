<template>
<div style="text-align: center">
    <div style="float:left; width:20%" class="sectionDiv">
        <h3><b-badge>조회기간</b-badge></h3>({{minFromDate.split(" ")[0]}} ~ {{maxToDate.split(" ")[0]}})
        <div style="width:80%; margin-left:auto; margin-right:auto"> 
            <h3>FROM</h3>
            <div style="float:center;">
                <b-col sm="9"><b-form-input id="fromDate" type="date" v-model="fromDate"></b-form-input></b-col>
                {{fromDate}}
            </div>
            <h3 style="padding-top:70px">TO</h3>
            <div style="float:center;">
                <b-col sm="9"><b-form-input id="toDate" type="date" v-model="toDate"></b-form-input></b-col>
                {{toDate}}
            </div>
        </div>
        <!-- <h3 style="padding-top: 100px"><b-badge>키워드</b-badge></h3>
        <div style="width:80%; margin-left:auto; margin-right:auto"> 
        <b-form-input v-model="keyword"
                    type="text"
                    placeholder="Enter keyword"></b-form-input>
        </div> -->
        <!-- 검색필터 
            기간, 시간대, 사용자(on/off), 키워드-->
        <div style="float:right; margin-right:10px">
            <b-button variant="primary" v-on:click="search">
                검색
            </b-button>
        </div>
    </div>
    <div style="float:left; width:53%" class="sectionDiv">
        <date-chart v-if="fromDate != '' && toDate != ''" :fromDate="fromDate" :toDate="toDate" :searchOrder="changeSearch" ></date-chart>
         <!-- 기간별 히트수 (전체, 각 사용자별 line차트 혹은 line-area) 그래프
                사용자별 히트수(막대 그래프)-->
    </div>
    <div style="float:left; width:24%" class="chatlog_frame">
        <h2>카카오톡 채팅 내역</h2>
        <chat-log v-if="fromDate != '' && toDate != ''" :fromDate="fromDate" :toDate="toDate" :searchOrder="changeSearch"></chat-log>
        <!-- 채팅내용 카톡처럼 -->
    </div>
</div>
</template>
    
<script>
import chatLog from '@/components/dateList/chatLog'
import dateChart from '@/components/dateList/dateChart'
export default {
    name : 'dataListPage',
    components : {
        'chat-log' : chatLog,
        'date-chart' : dateChart
    },
    data() {
        return {
            fromDate : "",
            toDate : "",
            keyword : "",
            maxToDate : "",
            minFromDate : "",
            changeSearch : 0
        }
    },
    created() {
        this.$http.get(`${this.api_url}/date/range`)
            .then((result) => {
                this.minFromDate = result.data.aggregations.minFromDate.value_as_string;
                this.maxToDate = result.data.aggregations.maxToDate.value_as_string;
                this.fromDate = this.minFromDate.split(" ")[0];
                this.toDate = this.maxToDate.split(" ")[0];
            })
    },
    mounted() {
        // var currentDate = new Date();
        // this.toDate = currentDate.getFullYear() + '-' + (currentDate.getMonth() + 1) + '-' + currentDate.getDate();
    },
    methods : {
        search() {
            this.changeSearch++;
        }
    }
}
</script>

<style>
    .sectionDiv {
        border: 1px solid gray; border-radius: 2em;
        margin-left: 12px;
        margin-top: 10px;
        height: 800px;
    }
    .chatlog_frame{
        border: none;
        margin-left: 12px;
        margin-top: 10px;
        height: 800px;
    }
</style>