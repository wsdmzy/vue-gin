<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
        >
        <b-card title="登陆">
          <b-form>
           
            <b-form-group label="手机号">
              <b-form-input 
                v-model="$v.user.telephone.$model" 
                type="number" 
                :state="validateState('telephone')"
                placeholder="请输入的你手机号" >
              </b-form-input> 
              <b-form-invalid-feedback :state="validateState('telephone')">
                手机号不符合规范
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input v-model="$v.user.password.$model" type="password" :state="validateState('password')" required placeholder="请输入你的密码"></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码必须大于6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group >
              <b-button 
                variant="outline-primary" 
                block
                @click="login"
                >登陆</b-button>
            </b-form-group>
          </b-form>
        </b-card>
       </b-col>
    </b-row>
  </div>
</template>

<script>
import { required,  between, minLength} from 'vuelidate/lib/validators'

import customValidator  from '@/common/validator.js'

import { mapActions } from 'vuex'

export default {
  data() {
    return {
      user: {
        telephone: '',
        password: ''
      },
      
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        phone: customValidator.telephoneValidator
      },
      password: {
        required,
        minLength: minLength(6)
      }
    }
  },
  methods: {
    ...mapActions('userModule', {userLogin: 'login'}),
    login() {
       // 验证数据
      this.$v.user.$touch(); //不填写表单提交也触发错误
      if (this.$v.user.$anyError) {
        return
      }
      // 请求
      this.userLogin(this.user).then(() => {
         // 跳转到主页
          this.$router.replace({ name: 'Home' })
      }).catch(err => {
        // console.log('err:' + err.response.data.msg)
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true 
        })
      })
    },
    validateState(name) {
      // $dirty交互了 
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null
    }
  }
};
</script>

<style>
</style>