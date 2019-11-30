import axios from "axios"

function success(app, callback) {
  return res => {
    if(res.status === 200 && res.data.code === 1) {
      app.$message.success("请求成功");
      callback(res.data.data);
    } else {
      app.$message.error("请求失败 " + res.data.message);
    }
  }
}

function error(app) {
  return err => {
    app.$message.error("请求失败 " + err);
  }
}

function ping(app, callback) {
  axios.get('http://localhost/api/ping').then(success(app, callback)).catch(error(app))
}

function getProblems(app, limit, offset, callback) {
  axios.get("http://localhost/api/v1/problem", {params: {limit: limit, offset: offset}})
    .then(success(app, callback)).catch(error(app))
}

function addLeetcodeProblem(app, url, callback) {
  if(url === "") {
    app.$message.warn("url为空");
    return
  }
  axios.post("http://localhost/api/v1/problem/url", {url: url})
    .then(success(app, callback)).catch(error(app))
}

function addProblem(app, problem, callback) {
  if(problem.content === "") {
    app.$message.error("content为空");
    return
  }
  axios.post("http://localhost/api/v1/problem", problem)
    .then(success(app, callback)).catch(error(app))
}

function getProblem(app, id, callback) {
  axios.get("http://localhost/api/v1/problem/" + id)
    .then(success(app, callback)).catch(error(app))
}

function editProblem(app, problem, callback) {
  axios.put("http://localhost/api/v1/problem", problem)
    .then(success(app, callback)).catch(error(app))
}

function delProblem(app, id, callback) {
  axios.delete("http://localhost/api/v1/problem", {data: {id: id}})
    .then(success(app, callback)).catch(error(app))
}

function getSolutions(app, problem_id, limit, offset, callback) {
  axios.get("http://localhost/api/v1/solution", {params: {problem_id: problem_id, limit: limit, offset: offset}})
    .then(success(app, callback)).catch(error(app))
}

function getSolution(app, id, callback) {
  axios.get("http://localhost/api/v1/solution/" + id)
    .then(success(app, callback)).catch(error(app))
}

function addSolution(app, solution, callback) {
  axios.post("http://localhost/api/v1/solution", solution)
    .then(success(app, callback)).catch(error(app))
}

function editSolution(app, solution, callback) {
  axios.put("http://localhost/api/v1/solution", solution)
    .then(success(app, callback)).catch(error(app))
}

function delSolution(app, id, callback) {
  axios.delete("http://localhost/api/v1/solution", {data: {id: id}})
    .then(success(app, callback)).catch(error(app))
}

export default {
  install: function (Vue) {
    Vue.prototype.http = {
      ping,
      
      getProblems,
      getProblem,
      addLeetcodeProblem,
      addProblem,
      editProblem,
      delProblem,
      
      getSolutions,
      getSolution,
      addSolution,
      editSolution,
      delSolution,
    };
  }
}