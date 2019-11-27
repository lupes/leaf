<template>
    <div>
        <a-row>
            <a-col :offset="19" :span="2">
                <a-button type="primary" @click="showAddProblem">添加题目</a-button>
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
        <a-row>
            <a-col :offset="4" :span="16">
                <a-list itemLayout="horizontal" :pagination="pagination"
                        :dataSource="problems">
                    <a-list-item slot="renderItem" slot-scope="item">
                        <span style="margin-right: 10px; width: 100px; ">{{ item.title }}</span>
                        <span>{{ item.created_time | datetime }}</span>
                        <a slot="actions" @click="showEditProblem(item)">编辑</a>
                        <a slot="actions" :href="'/problem/'+item.id">详情</a>
                        <a slot="actions" v-if="item.url !== ''" :href="item.url">链接</a>
                        <a slot="actions" @click="delProblem(item)">删除</a>
                    </a-list-item>
                </a-list>
            </a-col>
        </a-row>
    </div>
</template>

<script>
  import axios from "axios"

  export default {
    name: "Problems",
    data() {
      return {
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
        this.problem.title = problem.title;
        this.problem.url = problem.url;
        this.problem.content = problem.content;
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
          console.warn(problem);
          return
        }
        axios.post("http://localhost/api/v1/problem", problem).then(function (res) {
          app.$message.success("添加成功");
          app.$router.push("/problem/"+res.data.data)
          // app.getProblems(app.pageSize * (app.page - 1), res => {
          //   app.pagination.total = res.data.data.count;
          //   app.problems = res.data.data.problems == null ? [] : res.data.data.problems;
          // });
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
