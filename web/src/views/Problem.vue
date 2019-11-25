<template>
    <div>
        <a-row>
            <a-col :offset="19" :span="2">
                <a-button>Add</a-button>
            </a-col>
        </a-row>
        <a-row>
            <a-col :offset="4" :span="16">
                <a-list class="demo-loadmore-list" itemLayout="horizontal" :pagination="pagination"
                        :dataSource="problems">
                    <a-list-item slot="renderItem" slot-scope="item">
                        <a slot="actions" :href="'/problem/'+item.id">{{ item.title }}</a>
                        <a slot="actions">Edit</a>
                        <a slot="actions">Detail</a>
                        <a slot="actions">Solutions</a>
                        <a-list-item-meta :description="item.content" />
                    </a-list-item>
                </a-list>
            </a-col>
        </a-row>
    </div>
</template>

<script>
  import axios from "axios"

  export default {
    name: "problems",
    data() {
      return {
        problems: [],
        msg: "",
        pageSize: 10,
        pagination: {
          onChange: page => {
            this.getData(this.pageSize * (page - 1), res => {
              this.pagination.total = res.data.data.count;
              this.problems = res.data.data.problems == null ? [] : res.data.data.problems;
            });
          },
          total:2,
          pageSize: 10,
        },
      };
    },
    mounted() {
      this.getData(0, res => {
        console.log(res.data.data);
        this.pagination.total = res.data.data.count;
        this.problems = res.data.data.problems == null ? [] : res.data.data.problems;
      });
    },
    methods: {
      getData(offset, callback) {
        let app = this;
        axios.get("http://localhost/api/v1/problem", {
          params: {limit: app.pageSize, offset: offset}
        }).then(callback)
          .catch(function (error) {
            console.log('Error! Could not reach the API. ' + error)
            app.msg = 'Error! Could not reach the API. ' + error
          })
      },
    },
  };
</script>
<style>
    .demo-loadmore-list {
        min-height: 350px;
        text-align: center;
    }

    .home {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;

        color: #2c3e50;
    }
</style>
