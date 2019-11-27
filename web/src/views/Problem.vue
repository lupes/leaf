<template>
    <div>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <Back />
                <a-button type="primary" :href="getAddSolutionUrl" size="small" style="float: right;">添加题解</a-button>
            </a-col>
        </a-row>
        <a-row  type="flex" justify="center" style="margin: 20px 0">
            <a-col :span="18">
                <span><strong><a :href="problem.url">{{ problem.title }}</a></strong></span>
                <span style="float: right;">{{ problem.created_time | datetime }} </span>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <p v-html="problem.content"></p>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <a-list class="" itemLayout="horizontal" :pagination="pagination" header="题解列表"
                        :dataSource="solutions">
                    <a-list-item slot="renderItem" slot-scope="item">
                        <span>{{ item.title }}</span>
                        <span slot="actions">{{ item.created_time | datetime }}</span>
                        <a slot="actions" :href="'/solution/edit/'+item.id">编辑</a>
                        <a slot="actions" :href="'/solution/detail/'+item.id">详情</a>
                    </a-list-item>
                </a-list>
            </a-col>
        </a-row>
    </div>
</template>

<script>
  import axios from "axios"

  export default {
    name: "Problem",
    data() {
      return {
        problem: {},
        solutions: [],
        pageSize: 10,
        pagination: {
          onChange: page => {
            this.getSolutions(this.pageSize * (page - 1), this.$route.params.id, res => {
              this.pagination.total = res.data.data.count;
              this.solutions = res.data.data.solutions == null ? [] : res.data.data.solutions;
            });
          },
          total: 0,
          pageSize: 10,
        },
      };
    },
    mounted() {
      this.getProblem(this.$route.params.id, res => {
        this.problem = res.data.data
      });
      this.getSolutions(0, this.$route.params.id, res => {
        this.pagination.total = res.data.data.count;
        this.solutions = res.data.data.solutions == null ? [] : res.data.data.solutions;
      })
    },
    computed: {
      getAddSolutionUrl() {
        return "/solution/add/" + this.$route.params.id
      },
    },
    methods: {
      getProblem(problem_id, callback) {
        let app = this;
        axios.get("http://localhost/api/v1/problem/" + problem_id)
          .then(callback)
          .catch(function (error) {
            app.$message.error("请求失败 " + error);
          })
      },
      getSolutions(offset, problem_id, callback) {
        let app = this;
        axios.get("http://localhost/api/v1/solution", {
          params: {problem_id: problem_id, limit: this.pageSize, offset: offset}
        }).then(callback).catch(function (error) {
          app.$message.error("请求失败 " + error);
        })
      },
    },
  };
</script>

<style scoped>

</style>