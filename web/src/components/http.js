import axios from "axios"

function wrap(app, res, callback) {
  if(res.status === 200 && res.data.code === 1) {
    app.$message.success("请求成功");
    callback(res.data.data);
  } else {
    app.$message.error("请求失败 " + res.data.message);
  }
}

export default {
  install: function (Vue) {
    Vue.prototype.httpGetProblems = (app, limit, offset, callback) => {
      axios.get("http://localhost/api/v1/problem", {
        params: {limit: limit, offset: offset}
      }).then(res => {
        wrap(app, res, callback)
      }).catch(err => {
        app.$message.error("请求失败 " + err);
      })
    };
    Vue.prototype.httpAddLeetcodeProblem = (app, url, callback) => {
      if(url === "") {
        app.$message.warn("url为空");
        return
      }
      axios.post("http://localhost/api/v1/problem/url", {url: url}).then(res => {
        wrap(app, res, callback);
      }).catch(err => {
        app.$message.error("请求失败 " + err);
      });
    };
    Vue.prototype.httpAddProblem = (app, problem, callback) => {
      if(problem.content === "") {
        app.$message.warn("content为空");
      }
      axios.post("http://localhost/api/v1/problem", problem).then(res => {
        wrap(app, res, callback)
      }).catch(err => {
        app.$message.error("请求失败 " + err);
      });
    };
    
    Vue.prototype.httpEditProblem = (app, problem, callback) => {
      axios.put("http://localhost/api/v1/problem", problem).then(res => {
        wrap(app, res, callback)
      }).catch(err => {
        app.$message.error("请求失败 " + err);
      });
    };
    Vue.prototype.httpDelProblem = (app, id, callback) => {
      axios.delete("http://localhost/api/v1/problem", {data: {id: id}}).then(res => {
        wrap(app, res, callback)
      }).catch(err => {
        app.$message.error("请求失败 " + err);
      });
    }
  }
}