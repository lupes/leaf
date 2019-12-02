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
                <a-table :columns="columns" :dataSource="[problem]" :pagination="false" rowKey="id" :showHeader="false">
                    <a slot="name" slot-scope="item" :href="item.url">{{ item.title }}</a>
                    <a-tag slot="difficulty" :color="difficulty(item)" slot-scope="item">
                        {{ difficultyTransfer(item)}}
                    </a-tag>
                    <span slot="topics" slot-scope="topics" v-if="topics!==''">
                        <a-tag v-for="topic in topics.split(';')" :key="topic">{{topic}}</a-tag>
                    </span>
                    <span slot="created_time" slot-scope="item" style="padding-right: 0;">{{ item | datetime }}</span>
                </a-table>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 30px 0">
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
        columns: [
          {
            title: '题目',
            key: 'title',
            scopedSlots: {customRender: 'name'},
          },
          {
            title: '难度',
            dataIndex: 'difficulty',
            scopedSlots: {customRender: 'difficulty'},
          },
          {
            title: '标签',
            dataIndex: 'topics',
            scopedSlots: {customRender: 'topics'},
          },
          {
            title: '创建时间',
            dataIndex: 'created_time',
            align: 'right',
            scopedSlots: {customRender: 'created_time'},
          }
        ],
        problem: {
          id: this.$route.params.id,
          title: "",
          difficulty: "",
          topics: "",
          url: "",
          content: "",
        },
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