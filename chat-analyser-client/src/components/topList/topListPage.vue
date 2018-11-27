<template>
<!-- <div style="padding:2% 1px 1px 10%;"> -->
<div style="padding-left:25%;padding-right:1">
    <div style="width:35%; float:left">
        <b-table hover :fields="colum" :items="resultDatas.buckets"
                :sort-by.sync="sortBy"
                :sort-desc.sync="sortDesc"
                :current-page="currentPage"
                :per-page="10"
                @row-clicked="rowClick"
                caption-top
                >
            <template slot="table-caption">
                <p style="text-align:center;">키워드 TOP100</p>
            </template>
            <template slot="index" slot-scope="data">
                {{data.item.index + 1}}
            </template>
            <template slot="key" slot-scope="data">
                {{data.item.key}}
            </template>
            <template slot="doc_count" slot-scope="data">
                {{data.item.doc_count}}
            </template>
        
        </b-table>
        <b-pagination-nav base-url="#" :number-of-pages="10" v-model="currentPage" align="center"/>
        <p>
        Sorting By: <b>{{ sortBy }}</b>,
        Sort Direction: <b>{{ sortDesc ? 'Descending' : 'Ascending' }}</b>
        </p>
    </div>
    <div style="width:30%; float:left; margin-left:10px" v-cloak>
        <div v-if="resultDatas.buckets != null">
            <b-table hover :fields="authorColum" :items="resultDatas.buckets[currentSelectIdx].author.buckets"
                    caption-top
                    >
                <template slot="table-caption">
                    <p style="text-align:center;"><strong>{{resultDatas.buckets[currentSelectIdx].key}} </strong>사용현황</p>
                </template>
                <template slot="key" slot-scope="data">
                    {{data.item.key}}
                </template>
                <template slot="doc_count" slot-scope="data">
                    {{data.item.doc_count}}
                </template>
            
            </b-table>
            <pie-chart :data="chartData" :options="options"></pie-chart>
        </div>
            <!-- <div ref="piChart"></div> -->
    </div>
</div>
</template>

<script>
// TODO : 사용자별 키워드 분포 그래프 추가
// import tuiChart from 'tui-chart'
import 'tui-chart/dist/tui-chart.css'
import {pieChart} from '@toast-ui/vue-chart'
export default {
    name : 'topListPage',
    components : {
        'pie-chart' : pieChart
    },
    data() {
        return {
            resultDatas : {},
            currentKeywords :[],
            sortBy: 'index',
            sortDesc: false,
            colum : [
                {key:"index", label:"순위", sortable:true},
                {key:"key", label:"키워드", sortable:true},
                {key:"doc_count", label:"hit", sortable:true}
            ],
            authorColum : [
                {key:"key", label:"사용자", sortable:true},
                {key:"doc_count", label:"hit", sortable:true}
            ],
            currentSelectIdx : 0,
            currentPage : 1,
            options : {
                chart: {
                    title: '키워드별 사용자 분포'
                },
                tooltip: {
                    suffix: '%'
                },
                series: {
                    showLegend: true,
                    showLabel: true,
                    labelAlign: 'center'
                },
                legend: {
                    visible: true
                }
            },
            chartData : {
                categories: ['사용자'],
                series: []
            }
        }
    },
    methods : {
        rowClick(obj, idx) {
            this.currentSelectIdx = obj.index;
            //this.drawChart();
            this.chartData.series.splice(0,this.chartData.series.length)
            for(let author of this.resultDatas.buckets[this.currentSelectIdx].author.buckets) {
                    this.chartData.series.push({name:author.key, data:author.doc_count});
                }
        }
    },
    created() {
        this.$http.get(`${this.api_url}/top`, {params:{size:10, from:(this.currentPage-1)*10}, headers:{}, timeout:10000})
            .then((result) => {
                console.log(result)
                this.resultDatas = result.data.aggregations.keywords
                for(var i = 0; i<this.resultDatas.buckets.length; i++) {
                    this.resultDatas.buckets[i].index = i;
                }

                for(let author of this.resultDatas.buckets[this.currentSelectIdx].author.buckets) {
                    this.chartData.series.push({name:author.key, data:author.doc_count});
                }
                
            })
    },
    mounted() {
        this.$nextTick(() => {
            // this.drawChart();
            // this.$forceUpdate();
        })
        
    }
}
</script>


<style>
</style>