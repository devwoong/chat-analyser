<template>
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
</template>

<script>
export default {
    data() {
        return {
            currentKeywords :[],
            sortBy: 'index',
            sortDesc: false,
            colum : [
                {key:"index", label:"순위", sortable:true},
                {key:"key", label:"키워드", sortable:true},
                {key:"doc_count", label:"hit", sortable:true}
            ],
            currentSelectIdx : 0,
            currentPage : 1
        }
    },
    methods : {
        rowClick(obj, idx) {
            this.currentSelectIdx = obj.index;
        }
    },
    created() {
    },
    mounted() {

    },
    props : [
        'resultDatas']
}
</script>