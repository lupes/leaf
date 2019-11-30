<template>
    <div>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <Back/>
                <a-dropdown style="float: right;">
                    <a-menu slot="overlay">
                        <a-menu-item key="1" @click="()=> this.open = true">Leetcode</a-menu-item>
                        <a-menu-item key="2" @click="showAddProblem">手动添加</a-menu-item>
                    </a-menu>
                    <a-button> 添加题目
                        <a-icon type="down"/>
                    </a-button>
                </a-dropdown>
                <a-modal
                        title="添加Leetcode题目"
                        cancelText="取消"
                        okText="确认"
                        centered
                        :closable="false"
                        v-model="open"
                        @ok="handleLeetcodeOk">
                    <div style="margin: 10px 0"></div>
                    <a-input placeholder="题目连接" v-model="url"/>
                </a-modal>
                <problem-modal @ok="handlerOk"
                               @cancel="handlerCancel"
                               @change="closeProblem"
                               :modelTitle="title"
                               :problemTitle="problem.title"
                               :problemUrl="problem.url"
                               :problemContent="problem.content"
                               v-model="visible"/>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <a-list itemLayout="horizontal" :pagination="pagination" header="题解列表"
                        :dataSource="problems">
                    <a-list-item slot="renderItem" slot-scope="item">
                        <span>{{ item.title }}</span>
                        <span slot="actions">{{ item.created_time | datetime }}</span>
                        <a slot="actions" @click="showEditProblem(item)">编辑</a>
                        <a slot="actions" :href="'/problem/'+item.id">详情</a>
                        <a slot="actions" v-if="item.url !== ''" :href="item.url">链接</a>
                        <a-popconfirm title="确认删除?" @confirm="delProblem(item)" okText="是" cancelText="否"
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
    name: "Problems",
    data() {
      return {
        open: false,
        url: "",
        title: "",
        problem: {
          title: "",
          url: "",
          content: "",
        },
        problems: [],
        visible: false,
        pageSize: 10,
        page: 1,
        pagination: {
          onChange: page => {
            this.getProblems(page, this.pageSize);
          },
          total: 0,
          pageSize: 10,
        },
      };
    },
    mounted() {
      this.getProblems(1, this.pageSize)
    },
    methods: {
      closeProblem(f) {
        this.visible = f
      },
      handlerOk(problem) {
        console.warn(problem)
      },
      handlerCancel() {
        this.visible = false
      },
      showAddProblem() {
        this.problem.title = "";
        this.problem.url = "";
        this.problem.content = "";
        this.handlerOk = this.addProblem;
        this.title = "新建题目";
        this.visible = true;
      },
      showEditProblem(problem) {
        this.problem = problem;
        this.handlerOk = this.editProblem(problem.id);
        this.visible = true;
        this.title = "编辑题目";
      },
      getProblems(page, pageSize) {
        this.http.getProblems(this, pageSize, pageSize * (page - 1), data => {
          this.pagination.total = data.count;
          this.problems = data.problems == null ? [] : data.problems;
        });
      },
      handleLeetcodeOk() {
        this.open = false;
        this.http.addLeetcodeProblem(this, this.url, data => {
          this.$router.push("/problem/" + data)
        });
      },
      addProblem() {
        return this.http.addProblem(this, this.problem, data => {
          this.$router.push("/problem/" + data)
        })
      },
      editProblem(id) {
        return (problem) => {
          this.visible = false;
          problem.id = id;
          this.http.editProblem(this, problem, () => {
            this.getProblems(this.page, this.pageSize)
          });
        }
      },
      delProblem(item) {
        this.http.delProblem(this, item.id, () => {
          this.getProblems(this.page, this.pageSize)
        })
      },
    },
  }
</script>

<style>

</style>
