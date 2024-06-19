<template>
  <CommonPage>
    <template #action>
      <NButton type="primary" @click="handleAdd()">
        <i class="i-material-symbols:add mr-4 text-18" />
        创建新杆塔
      </NButton>
    </template>
    <MeCrud ref="$table" v-model:query-items="queryItems" :scroll-x="2000" :columns="columns" :get-data="api.read">
      <MeQueryItem label="项目包" :label-width="50">
        <n-select v-model:value="queryItems.project_id" clearable :options="statusOptions" />
      </MeQueryItem>

      <MeQueryItem label="子项" :label-width="50">
        <n-select v-model:value="queryItems.subitem_id" clearable :options="statusOptions" />
      </MeQueryItem>
      <!-- <MeQueryItem label="性别" :label-width="50">
        <n-select v-model:value="queryItems.gender" clearable :options="genders" />
      </MeQueryItem>

      <MeQueryItem label="状态" :label-width="50">
        <n-select v-model:value="queryItems.enable" clearable :options="[
          { label: '启用', value: 1 },
          { label: '停用', value: 0 },
        ]" />
      </MeQueryItem> -->
    </MeCrud>
    <!-- 这里是编辑表单的填写 -->
    <MeModal ref="modalRef" width="520px">
      <n-form ref="modalFormRef" label-placement="left" label-align="left" :label-width="80" :model="modalForm"
        :disabled="modalAction === 'view'">
        <n-form-item label="用户名" path="username" :rule="{
          required: true,
          message: '请输入用户名',
          trigger: ['input', 'blur'],
        }">
          <n-input v-model:value="modalForm.username" :disabled="modalAction !== 'add'" />
        </n-form-item>

        <!-- <n-form-item label="项目包">
          <n-input v-model:value="modalForm.project_id" placeholder="请输入项目包ID" />
        </n-form-item> -->
        <n-form-item label="子项名称">
          <n-input v-model:value="modalForm.subitem_id" placeholder="请输入子项名称" />
        </n-form-item>
        <n-form-item label="杆塔名称">
          <n-input v-model:value="modalForm.name" placeholder="请输入杆塔名称" />
        </n-form-item>
        <n-form-item label="详细地区">
          <n-input v-model:value="modalForm.address" placeholder="请输入详细地区" />
        </n-form-item>
        <!-- <n-form-item label="基站图片">
          <n-upload v-model:value="modalForm.image" />
        </n-form-item> -->
        <n-form-item label="检查状态">
          <n-select v-model:value="modalForm.check_status" :options="statusOptions" />
        </n-form-item>
        <n-form-item label="检查时间"> 
          <n-input v-model:value="modalForm.check_time"  placeholder="选择检查时间" />
        </n-form-item> 
        <n-form-item label="检查人">
          <n-input v-model:value="modalForm.check_user" placeholder="请输入检查人姓名" />
        </n-form-item>
        <n-form-item label="负责人">
          <n-input v-model:value="modalForm.principal_id" placeholder="请输入负责人ID" />
        </n-form-item>
        <n-form-item label="计划时间">
          <n-input v-model:value="modalForm.plan_check_time"  placeholder="选择计划时间" />
        </n-form-item>

        <n-form-item v-if="['add', 'reset'].includes(modalAction)" :label="modalAction === 'reset' ? '重置密码' : '初始密码'"
          path="password" :rule="{
            required: true,
            message: '请输入密码',
            trigger: ['input', 'blur'],
          }">
          <n-input v-model:value="modalForm.password" />
        </n-form-item>

        <n-form-item v-if="['add', 'setRole'].includes(modalAction)" label="角色" path="roleIds">
          <n-select v-model:value="modalForm.roleIds" :options="roles" label-field="name" value-field="id" clearable
            filterable multiple />
        </n-form-item>
        <n-form-item v-if="modalAction === 'add'" label="状态" path="enable">
          <NSwitch v-model:value="modalForm.enable">
            <template #checked>
              启用
            </template>
            <template #unchecked>
              停用
            </template>
          </NSwitch>
        </n-form-item>
      </n-form>
      <n-alert v-if="modalAction === 'add'" type="warning" closable>
        详细信息需由用户本人补充修改
      </n-alert>
    </MeModal>
  </CommonPage>
</template>

<script setup>
import { NAvatar, NButton, NSwitch, NTag } from 'naive-ui'
import api from './api'
import { formatDateTime } from '@/utils'
import { MeCrud, MeModal, MeQueryItem } from '@/components'
import { useCrud } from '@/composables'

defineOptions({ name: 'UserMgt' })

const $table = ref(null)
/** QueryBar筛选参数（可选） */
const queryItems = ref({})

onMounted(() => {
  $table.value?.handleSearch()
})

const statusOptions = ref([
  { label: '启用', value: 1 },
  { label: '停用', value: 0 },
  { label: 'test', value: 4 },
])


const genders = [
  { label: '男', value: 1 },
  { label: '女', value: 2 },
]
const roles = ref([])
api.getAllRoles().then(({ data = [] }) => (roles.value = data))

const {
  modalRef,
  modalFormRef,
  modalForm,
  modalAction,
  handleAdd,
  handleDelete,
  handleOpen,
  handleSave,
  handleEdit,
} = useCrud({
  name: '杆塔',
  initForm: { enable: true },
  doCreate: api.create,
  doDelete: api.delete,
  doUpdate: api.update,
  refresh: () => $table.value?.handleSearch(),
})
const statusMap = {
  "1": '未完成',
  "2": '已完成',
  // 添加更多状态映射
};

const columns = [
  { title: 'id', key: 'id', width: 50, ellipsis: { tooltip: true } },
  { title: '项目包', key: 'project_id', width: 80, ellipsis: { tooltip: true } },
  { title: '子项名称', key: 'subitem_id', width: 80, ellipsis: { tooltip: true } },
  { title: '杆塔名称', key: 'name', width: 100, ellipsis: { tooltip: true } },
  { title: '详细地区', key: 'address', width: 120, ellipsis: { tooltip: true } },
  {
    title: '基站图片',
    key: 'image',
    width: 80,
    render: ({ image }) =>
      h(NAvatar, {
        size: 'medium',
        src: image.length > 0 ? image[0] : '', // 使用数组的第一个值作为图片源
      }),
  },
  {
    title: '检查状态',
    key: 'check_status',
    width: 80,
    render: ({ check_status }) =>
      h(NTag, {
        type: check_status == 1 ? 'warning' : 'success',
        // 根据状态值显示对应的文本
      }, { default: () => statusMap[check_status] }),
  },
  { title: '检查时间', key: 'check_time', width: 150, ellipsis: { tooltip: true } },
  { title: '检查人', key: 'check_user', width: 50, ellipsis: { tooltip: true } },
  { title: '负责人', key: 'principal_id', width: 50, ellipsis: { tooltip: true } },
  { title: '计划时间', key: '计划检查时间', width: 150, ellipsis: { tooltip: true } },
  // {
  //   title: '性别', // 表格列的标题
  //   key: 'gender', // 对应列内容的字段名，用于从数据对象中取值
  //   width: 80, // 设置列的宽度为80像素
  //   render: ({ gender }) => genders.find(item => gender === item.value)?.label ?? '2', // 自定义渲染函数，用于显示列的内容
  //   // render函数接收一个参数，其中包含了gender字段的值
  //   // 使用genders数组查找与gender字段值相匹配的项，并返回该项的label属性作为显示内容
  //   // 如果没有找到匹配项，则显示空字符串
  // },
  // {
  //   title: '创建时间',
  //   key: 'createDate',
  //   width: 180,
  //   render(row) {
  //     return h('span', formatDateTime(row.createTime))
  //   },
  // },
  // {
  //   title: '状态',
  //   key: 'enable',
  //   width: 120,
  //   render: row =>
  //     h(
  //       NSwitch,
  //       {
  //         size: 'small',
  //         rubberBand: false,
  //         value: row.enable,
  //         loading: !!row.enableLoading,
  //         onUpdateValue: () => handleEnable(row),
  //       },
  //       {
  //         checked: () => '启用',
  //         unchecked: () => '停用',
  //       },
  //     ),
  // },
  {
    title: '操作',
    key: 'actions',
    width: 140,
    align: 'right',
    fixed: 'right',
    hideInExcel: true,
    render(row) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            style: 'margin-left: 12px;',
            onClick: () => handleOpen({ action: 'reset', title: '查看', row, onOk: onSave }),
          },
          {
            default: () => '查看',
            icon: () => h('i', { class: 'i-radix-icons:reset text-14' }),
          },
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            style: 'margin-left: 12px;',
            disabled: row.code === 'SUPER_ADMIN',
            onClick: () => handleEdit(row),
          },
          {
            default: () => '编辑',
            icon: () => h('i', { class: 'i-material-symbols:edit-outline text-14' }),
          },
        ),

        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            style: 'margin-left: 12px;',
            onClick: () => handleOpen({ action: 'reset', title: '查看', row, onOk: onSave }),
          },
          {
            default: () => '编辑',
            icon: () => h('i', { class: 'i-radix-icons:reset text-14' }),
          },
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'error',
            style: 'margin-left: 12px;',
            onClick: () => handleDelete(row.id),
          },
          {
            default: () => '删除',
            icon: () => h('i', { class: 'i-material-symbols:delete-outline text-14' }),
          },
        ),
      ]
    },
  },
]

async function handleEnable(row) {
  row.enableLoading = true
  try {
    await api.update({ id: row.id, enable: !row.enable })
    row.enableLoading = false
    $message.success('操作成功')
    $table.value?.handleSearch()
  }
  catch (error) {
    row.enableLoading = false
  }
}

function handleOpenRolesSet(row) {
  const roleIds = row.roles.map(item => item.id)
  handleOpen({
    action: 'setRole',
    title: '分配角色',
    row: { id: row.id, username: row.username, roleIds },
    onOk: onSave,
  })
}

function onSave() {
  if (modalAction.value === 'setRole') {
    return handleSave({
      api: () => api.update(modalForm.value),
      cb: () => $message.success('分配成功'),
    })
  }
  else if (modalAction.value === 'reset') {
    return handleSave({
      api: () => api.resetPwd(modalForm.value.id, modalForm.value),
      cb: () => $message.success('密码重置成功'),
    })
  }
  handleSave()
}
</script>
