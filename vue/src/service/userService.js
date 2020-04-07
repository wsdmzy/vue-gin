import request from '@/common/request'

// 用户注册
const register = ({ name, telephone, password }) => {
  return request.post('auth/register', {name, telephone, password})
}

// 获取用户信息
const info = () => {
  return request.get('auth/info')
}

// 登陆
const login = ({telephone, password}) => {
  return request.post('auth/login', {telephone, password})
}

export default {
  register,
  info,
  login
}