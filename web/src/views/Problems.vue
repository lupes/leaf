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
                        <a-popconfirm
                                title="确认删除?"
                                @confirm="delProblem(item)"
                                okText="是"
                                cancelText="否"
                                slot="actions"
                        >
                            <a href="#">删除</a>
                        </a-popconfirm>
                    </a-list-item>
                </a-list>
            </a-col>
        </a-row>
    </div>
</template>

<script>
  import axios from "axios"
  import Back from "../components/Back";

  export default {
    name: "Problems",
    components: {Back},
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
            this.getProblems(this.pageSize * (page - 1), res => {
              this.page = page;
              this.pagination.total = res.data.data.count;
              this.problems = res.data.data.problems == null ? [] : res.data.data.problems;
            });
          },
          total: 0,
          pageSize: 10,
        },
      };
    },
    mounted() {
      this.getProblems(0, res => {
        this.pagination.total = res.data.data.count;
        this.problems = res.data.data.problems == null ? [] : res.data.data.problems;
      });
    },
    methods: {
      handleLeetcodeOk() {
        let app = this;
        this.open = false;
        if(app.url === "") {
          app.$message.warn("content为空");
        }
        axios.post("http://localhost/api/v1/problem/url", {url: app.url}).then(function (res) {
          app.$message.success("添加成功");
          app.$router.push("/problem/" + res.data.data)
        }).catch(function (error) {
          app.$message.error("请求失败 " + error);
        });
      },
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
      getProblems(offset, callback) {
        let app = this;
        axios.get("http://localhost/api/v1/problem", {
          params: {limit: this.pageSize, offset: offset}
        }).then(callback).catch(function (error) {
          app.$message.error("请求失败 " + error);
        })
      },
      addProblem(problem) {
        let app = this;
        this.visible = false;
        if(problem.content === "") {
          app.$message.warn("content为空");
        }
        axios.post("http://localhost/api/v1/problem", problem).then(function (res) {
          app.$message.success("添加成功");
          app.$router.push("/problem/" + res.data.data)
        }).catch(function (error) {
          app.$message.error("请求失败 " + error);
        });
      },
      editProblem(id) {
        return (problem) => {
          let app = this;
          this.visible = false;
          axios.put("http://localhost/api/v1/problem", {
            problem_id: id,
            title: problem.title,
            url: problem.url,
            content: problem.content,
          }).then(function () {
            app.$message.success("编辑成功");
            app.getProblems(app.pageSize * (app.page - 1), res => {
              app.pagination.total = res.data.data.count;
              app.problems = res.data.data.problems == null ? [] : res.data.data.problems;
            });
          }).catch(function (error) {
            app.$message.error("请求失败 " + error);
          });
        }
      },
      delProblem(item) {
        let app = this;
        axios.delete("http://localhost/api/v1/problem", {
          data: {
            problem_id: item.id
          }
        }).then(function () {
          app.$message.success("删除成功");
          app.getProblems(app.pageSize * (app.page - 1), res => {
            app.pagination.total = res.data.data.count;
            app.problems = res.data.data.problems == null ? [] : res.data.data.problems;
          });
        }).catch(function (error) {
          app.$message.error("请求失败 " + error);
        });
      },
    },
  };
</script>

<style>

</style>
