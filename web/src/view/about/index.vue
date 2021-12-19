<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="24">
        <el-card>
          <template #header>
            <el-divider
              style="
                display: grid;
                justify-content: center;
                align-items: center;
                grid-auto-flow: column;
            "
            >
              <el-image
                style="width: 20px; height: 20px; border-radius: 50%; padding: 4px;"
                src="http://lumingguoji.com/logo_center.svg"
                fit="fit"
              />
              <span>MAXFUTURE</span></el-divider>
            鹿名国际教育成立于2013年美国波士顿，其使命致力于让中国留学生的名校梦不再高不可攀。
            不论学生是怎样的背景和成绩，鹿名国际都可以通过其两条成熟和独特的留学申请产品线（世界名校直通车和名校教授培养计划）将学生保送到最顶尖的名校，
            并在留学后的校园生活中为学生提供系统性的学术辅导服务（学业管家）。
            鹿名国际凭借着多年以来独有的世界顶尖大学教授、招生官、学校管理层以及名校学长学姐资源，为学生的申请路上保驾护航。
          </template>
        </el-card></el-col></el-row>
  </div>
</template>

<script>
import { Commits, Members } from '@/api/github'
export default {
  name: 'About',
  data() {
    return {
      messageWhenNoItems: 'There arent commits',
      members: [],
      dataTimeline: [],
      page: 0,
    }
  },
  created() {
    this.loadCommits()
    this.loadMembers()
  },
  methods: {
    loadMore() {
      this.page++
      this.loadCommits()
    },
    loadCommits() {
      Commits(this.page).then(({ data }) => {
        data.forEach((element) => {
          if (element.commit.message) {
            this.dataTimeline.push({
              from: new Date(element.commit.author.date),
              title: element.commit.author.name,
              showDayAndMonth: true,
              message: element.commit.message,
            })
          }
        })
      })
    },
    loadMembers() {
      Members().then(({ data }) => {
        this.members = data
        this.members.sort()
      })
    },
  },
}
</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
