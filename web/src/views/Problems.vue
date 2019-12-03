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
                               :modal-title="title"
                               :problem-no="problem.no"
                               :problem-title="problem.title"
                               :problem-difficulty="problem.difficulty"
                               :problem-topics="problem.topics"
                               :problem-url="problem.url"
                               :problem-content="problem.content"
                               v-model="visible"/>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <a-table :columns="columns" :dataSource="problems" :pagination="pagination" rowKey="id" header="题目列表">
                    <a slot="name" slot-scope="item" :href="'/problem/'+item.id">{{ item.title }}</a>
                    <a-tag slot="difficulty" :color="difficulty(item)" slot-scope="item">{{ difficultyTransfer(item) }}
                    </a-tag>
                    <span slot="topics" slot-scope="topics">
                        <span v-if="topics!==''">
                            <a-tag v-for="topic in topics.split(';')" :key="topic">{{topic}}</a-tag>
                        </span>
                    </span>
                    <span slot="created_time" slot-scope="item">{{ item | datetime }}</span>
                    <span slot="action" slot-scope="item">
                      <a @click="showEditProblem(item)">编辑</a>
                      <a-divider type="vertical"/>
                      <a slot="actions" v-if="item.url !== ''" :href="item.url">Leetcode</a>
                      <a-divider type="vertical"/>
                      <a-popconfirm title="确认删除?" @confirm="delProblem(item)" okText="是" cancelText="否" slot="actions">
                         <a href="#">删除</a>
                      </a-popconfirm>
                    </span>
                </a-table>
            </a-col>
        </a-row>
    </div>
</template>

<script>
  export default {
    name: "Problems",
    data() {
      return {
        columns: [
          {
            title: '编号',
            dataIndex: 'no',
            key: 'no',
          },
          {
            title: '题目',
            key: 'title',
            scopedSlots: {customRender: 'name'},
          },
          {
            title: '难度',
            dataIndex: 'difficulty',
            key: 'difficulty',
            scopedSlots: {customRender: 'difficulty'},
          },
          {
            title: '标签',
            dataIndex: 'topics',
            key: 'topics',
            scopedSlots: {customRender: 'topics'},
          },
          {
            title: '创建时间',
            dataIndex: 'created_time',
            key: 'created_time',
            scopedSlots: {customRender: 'created_time'},
          },
          {
            title: '操作',
            key: 'action',
            scopedSlots: {customRender: 'action'},
          },
        ],
        open: false,
        url: "",
        title: "",
        problem: {
          no: "",
          title: "",
          url: "",
          topics: "",
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
      difficulty(value) {
        if(value.toLowerCase() === 'easy') {
          return 'green'
        } else if(value.toLowerCase() === 'medium') {
          return 'orange'
        } else if(value.toLowerCase() === 'hard') {
          return 'red'
        } else {
          return '???' + value
        }
      },
      difficultyTransfer(value) {
        if(value.toLowerCase() === 'easy') {
          return '普通'
        } else if(value.toLowerCase() === 'medium') {
          return '中等'
        } else if(value.toLowerCase() === 'hard') {
          return '困难'
        } else {
          return '???' + value
        }
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
