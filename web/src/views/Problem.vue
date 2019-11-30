<template>
    <div>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <Back/>
                <a-button type="primary" :href="getAddSolutionUrl" size="small" style="float: right;">添加题解</a-button>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 20px 0">
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
                <a-list class="" itemLayout="horizontal" :pagination="pagination" header="题解列表" :dataSource="solutions">
                    <a-list-item slot="renderItem" slot-scope="item">
                        <span>{{ item.title }}</span>
                        <span slot="actions">{{ item.created_time | datetime }}</span>
                        <a slot="actions" :href="'/solution/edit/'+item.id">编辑</a>
                        <a slot="actions" :href="'/solution/detail/'+item.id">详情</a>
                        <a-popconfirm title="确认删除?" @confirm="delSolution(item)" okText="是" cancelText="否"
                                      slot="actions">
                            <a href="#">删除</a>
                        </a-popconfirm>
                    </a-list-item>
                </a-list>
            </a-col>
        </a-row>
    </div>
</template>

<script>

  export default {
    name: "Problem",
    data() {
      return {
        problem: {},
        solutions: [],
        pageSize: 10,
        page: 0,
        pagination: {
          onChange: this.pageChange,
          total: 0,
          pageSize: 10,
        },
      };
    },
    mounted() {
      this.http.getProblem(this, this.$route.params.id, data => {
        this.problem = data
      });
      this.pageChange(1);
    },
    computed: {
      getAddSolutionUrl() {
        return "/solution/add/" + this.$route.params.id
      },
    },
    methods: {
      pageChange(page) {
        this.http.getSolutions(this, this.$route.params.id, this.pageSize, this.pageSize * (page - 1), data => {
          this.page = page;
          this.pagination.total = data.count;
          this.solutions = data.solutions == null ? [] : data.solutions;
        })
      },
      delSolution(item) {
        this.http.delSolution(this, item.id, () => {
          this.pageChange(this.page)
        })
      },
    },
  };
</script>

<style scoped>

</style>