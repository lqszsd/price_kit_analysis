
<template>
  <div>
<el-form label-position="right" label-width="80px">

  <el-form-item label="账号">
       <el-col :span="11">
    <el-input v-model="account"></el-input>
    </el-col>
  </el-form-item>
  <el-form-item label="密码">
      <el-col :span="11">
    <el-input type="password" v-model="password"></el-input>
    </el-col>
  </el-form-item>
  <el-form-item>
       <el-col :span="11">
    <el-button type="primary" @click="onSubmit">登录</el-button>
    </el-col>
  </el-form-item>
</el-form>

  </div>
</template>
<script>
import axios from 'axios';

export default {
  name: "NewList",
  components: {
  },
  data: function () {
    return {
    account: '',
    password: '',
     
    };
  },
  methods: {
      onSubmit:function(){
          var that=this;
          axios.post(
              "http://127.0.0.1:8888/login",{account:this.account,password:this.password}
          ).then(data=>{
              console.log(data)
              if(data.data.code=="400"){
                  that.$message.error(data.data.errors)
              } else{
                    localStorage.setItem("token", data.data.token);
                    window.location.href="/login";
                    return;

              }
          }).catch(function(err){
              console.log(err)
          })
      },
  },
  mounted: function () {
  },
};
</script>