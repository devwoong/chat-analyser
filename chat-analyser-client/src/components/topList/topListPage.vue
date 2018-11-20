<template>
<div class="text-center">
    <div style="width:50%">

        <b-table hover :fields="colum" :items="resultDatas.buckets"
                :sort-by.sync="sortBy"
                :sort-desc.sync="sortDesc"
                caption-top
                >
            <template slot="table-caption">
                <p style="text-align:center;">키워드 TOP100</p>
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
        <p>
        Sorting By: <b>{{ sortBy }}</b>,
        Sort Direction: <b>{{ sortDesc ? 'Descending' : 'Ascending' }}</b>
        </p>
    </div>
</div>
</template>

<script>
export default {
    name : 'topListPage',
    data() {
        return {
            resultDatas : {},
            sortBy: 'index',
            sortDesc: false,
            colum : [
                {key:"index", label:"순위", sortable:true},
                {key:"key", label:"키워드", sortable:true},
                {key:"doc_count", label:"hit", sortable:true}
            ]
        }
    },
    methods : {
    },
    created() {
        this.$http.get(`${this.api_url}/top`)
            .then((result) => {
                console.log(result)
                this.resultDatas = result.data.aggregations.keywords
            })
    },
    mounted() {

    }
}
</script>


<style>
</style>