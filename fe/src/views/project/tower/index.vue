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
    </MeCrud>

    <!-- 这里是编辑表单的填写 -->
    <MeModal ref="modalRef" width="520px">
      <n-form ref="modalFormRef" label-placement="left" label-align="left" :label-width="80" :model="modalForm"
        :disabled="modalAction === 'view'">
        <n-form-item label="子项id">
          <n-input-number v-model:value="modalForm.subitem_id" placeholder="请输入子项名称" />
        </n-form-item>
        <n-form-item label="杆塔名称">
          <n-input v-model:value="modalForm.name" placeholder="请输入杆塔名称" />
        </n-form-item>
        <n-form-item label="详细地区">
          <n-input v-model:value="modalForm.address" placeholder="请输入详细地区" />
        </n-form-item>
        <n-form-item label="检查状态">
          <n-select v-model:value="modalForm.check_status" :options="statusOptions" />
        </n-form-item>
        <n-form-item label="检查时间">
          <n-input v-model:value="modalForm.check_time" placeholder="选择检查时间" />
        </n-form-item>
        <n-form-item label="检查人">
          <n-input-number v-model:value="modalForm.check_user" placeholder="请输入检查人姓名" />
        </n-form-item>
        <n-form-item label="负责人">
          <n-input-number v-model:value="modalForm.principal_id" placeholder="请输入负责人ID" />
        </n-form-item>
        <n-form-item label="计划时间">
          <n-date-picker v-model:value="modalForm.plan_check_time" type="datetime" placeholder="选择计划时间"
            :value-format="'yyyy-MM-dd HH:mm:ss'" />
        </n-form-item>
      </n-form>
    </MeModal>


    <!-- todo 这里看一下要怎么解决获取数据异常问题-->
    <MeModal ref="modalRef" width="520px">
      <MeCrud ref="$deviceListTable" v-model:query-items="deviceListTableQueryItems" :scroll-x="200"
        :columns="deviceListColumns" :get-data="api.readDeviceList">
      </MeCrud>
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

// 这里的正常的浏览表格的
const $table = ref(null)
/** QueryBar筛选参数（可选） */
const queryItems = ref({})


// 这里是查看设备的逻辑
// 这里的正常的浏览表格的
const $deviceListTable = ref(null)
/** QueryBar筛选参数（可选） */
const deviceListTableQueryItems = ref({})

// 页面加载就执行查询方法
onMounted(() => {
  $table.value?.handleSearch()
  // $deviceListTable.value?.handleSearch()
})


// 定义响应式变量来存储表格数据
const deviceListData = ref([]);


// 检查状态的展示
const statusOptions = ref([
  { label: '启用', value: 1 },
  { label: '停用', value: 0 },
  { label: '待检测', value: 2 },
  { label: 'todo', value: 3 },
])

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
  "0": 'todotest',
  "1": '未完成',
  "2": '已完成',
  "3": '待检测',
  "4": 'todo',
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
  {
    title: '检查时间',
    key: 'check_time',
    width: 150,
    ellipsis: { tooltip: true },
    render: ({ check_time }) => check_time ? new Date(check_time).toLocaleString() : ''
  },
  { title: '检查人', key: 'check_user', width: 50, ellipsis: { tooltip: true } },
  { title: '负责人', key: 'principal_id', width: 50, ellipsis: { tooltip: true } },
  { title: '计划时间', key: '计划检查时间', width: 150, ellipsis: { tooltip: true } },
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
            onClick: () => handleGetDeviceList(row),
          },
          {
            default: () => '查看设备',
            icon: () => h('i', { class: 'i-radix-icons:reset text-14' }),
          },
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            style: 'margin-left: 12px;',
            onClick: () => handleGetDeviceList(row),
          },
          {
            default: () => '比对',
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


// 定义 handleFetchData 函数，用于请求接口并渲染数据
async function handleGetDeviceList(row) {
  // 打开模态框
  handleOpen({ action: 'view', row: deviceListData });
  console.log('111', $deviceListTable.value);
  deviceListTableQueryItems.id = row.id
  console.log('row.id', row.id, deviceListTableQueryItems.id);
  setTimeout(() => {
    deviceListTableQueryItems.pageNo = row.id
    deviceListTableQueryItems.value.pageNo = row.id
    $deviceListTable.value?.handleSearchModel(row.id)
    // $deviceListTable.value?.handleSearch()

  }, 1000);
}

// 查看设备的列表的table
const deviceListColumns = [
  // { title: 'id', key: 'id', width: 10, ellipsis: { tooltip: true } },
  { title: '设备列表', key: 'name', width: 10, ellipsis: { tooltip: true } },]


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