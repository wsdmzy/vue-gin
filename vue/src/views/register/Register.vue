<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
        >
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input v-model="user.name" type="text" placeholder="请输入你的名称 (选填)"></b-form-input>
            </b-form-group>
            <b-form-group label="手机号">
              <b-form-input v-model="user.telephone" type="number" required placeholder="请输入的你手机号" @change="show"></b-form-input>
               <b-form-text 
                  id="password-help-block" 
                  text-variant="danger" 
                  v-if="showTelephoneValidate">
                手机号必须11位
              </b-form-text>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input v-model="user.password" type="password" required placeholder="请输入你的密码"></b-form-input>
            </b-form-group>
            <b-form-group >
              <b-button 
                variant="outline-primary" 
                block
                @click="register"
                >注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
       </b-col>
    </b-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: ''
      },
      showTelephoneValidate: false
    };
  },
  methods: {
    register() {
      
      
    },
    show() {
      this.user.telephone.length !== 11 ? this.showTelephoneValidate = true : this.showTelephoneValidate = false
    }
  }
};
</script>

<style>
</style>