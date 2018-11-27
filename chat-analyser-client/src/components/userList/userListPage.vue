<template>
<div>
    <div style = "float:left; width:50%" v-cloak>
        <div style="float:left;">
            <b-table hover :fields="colum" :items="resultDatas.buckets"
                    :sort-by.sync="sortBy"
                    :sort-desc.sync="sortDesc"
                    @row-clicked="rowClick"
                    caption-top
                    >
                <template slot="table-caption">
                    <p style="text-align:center;">사용자순위</p>
                </template>
                <template slot="index" slot-scope="data">
                    {{data.index + 1}}
                </template>
                <template slot="key" slot-scope="data">
                    {{data.item.key}}
                </template>
                <template slot="doc_count" slot-scope="data">
                    {{data.item.doc_count}}
                </template>
            </b-table>
        </div>
        <div v-if="resultDatas.buckets != null" style="float:left;">
            <pie-chart :data="userChart.chartData" :options="userChart.options"></pie-chart>
        </div>
    </div>
    <div style = "float:left; width:50%">
        <b-tabs v-if="resultDatas.buckets != null" v-model="tabIndex">
            <b-tab :title="data.key + '(' + data.doc_count + ')'" v-for="(data, idx) in resultDatas.buckets" :key="idx">              
                <!-- <div style="width:35%; float:left">
                    <b-table hover :fields="colum" :items="userKeywords.buckets"
                            :sort-by.sync="sortBy"
                            :sort-desc.sync="sortDesc"
                            :current-page="currentPage"
                            :per-page="10"
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
                </div> -->
                <user-keywords :resultDatas="userKeywords"></user-keywords>
            </b-tab>
        </b-tabs>
    </div>
</div>
</template>

<script>
//  TODO : 분포도그래프 추가, 가장 많이 사용한 키워드,  

import userKeywordList from '@/components/userList/userKeywordList'
import {pieChart} from '@toast-ui/vue-chart'

export default {
    name : 'userListPage',
    components : {
        'user-keywords' : userKeywordList,
        'pie-chart' : pieChart
    },
    data() {
        return {
            resultDatas : {},
            userKeywords : {},
            sortBy: 'index',
            sortDesc: false,
            tabIndex : 0,
            colum : [
                {key:"index", label:"순위", sortable:true},
                {key:"key", label:"사용자", sortable:true},
                {key:"doc_count", label:"hit", sortable:true}
            ],
            userChart : {
                options : {
                    chart: {
                        title: '사용자 분포',
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
        }
    },
    created() {
        this.$http.get(`${this.api_url}/user`)
            .then((result) => {
                console.log(result)
                this.resultDatas = result.data.aggregations.author
                for(var i = 0; i<this.resultDatas.buckets.length; i++) {
                    this.userChart.chartData.series.push({
                        name : this.resultDatas.buckets[i].key,
                        data : this.resultDatas.buckets[i].doc_count})
                }
                // for(var i = 0; i<this.resultDatas.buckets.length; i++) {
                //     this.resultDatas.buckets[i].index = i;
                // }
            this.$http.get(`${this.api_url}/user/keywords/${this.resultDatas.buckets[this.tabIndex].key}`)
                .then((result) => {
                    console.log(result)
                    this.userKeywords = result.data.aggregations.keywords
                    for(var i = 0; i<this.userKeywords.buckets.length; i++) {
                        this.userKeywords.buckets[i].index = i;
                    }

                })
            })
        
        
    },
    methods : {
        rowClick(obj, idx) {
            this.tabIndex = idx;
        }

    },
    watch : {
        tabIndex() {

        this.$http.get(`${this.api_url}/user/keywords/${this.resultDatas.buckets[this.tabIndex].key}`)
            .then((result) => {
                console.log(result)
                this.userKeywords = result.data.aggregations.keywords
                for(var i = 0; i<this.userKeywords.buckets.length; i++) {
                    this.userKeywords.buckets[i].index = i;
                }
            })
        }
    }
}
</script>

<style>
</style>