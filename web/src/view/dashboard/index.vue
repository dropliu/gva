<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
        
          <div class="gva-top-card-left-rows">
            <el-row>
             
              <div class="card-container">
                <h3 style="text-align: left;">今日</h3>
              <el-card class="box-card">
                <div class="item">
                  <span>交易金额：</span>
                  <p>{{ staticData.amountDay }}</p>
                </div>
                <div class="item">
                  <span>佣金：</span>
                  <p>{{ staticData.commissionDay }}</p>
                </div>
                <div class="item">
                  <span>交易笔数：</span>
                  <p>{{ staticData.countDay }}</p>
                </div>
              </el-card>
   

              <el-card class="box-card">
                <div class="item">
                  <span>实时余额：</span>
                  <p>{{ staticData.amount }}</p>
                </div>
                <div class="item">
                  <span>待结算：</span>
                  <p>{{ staticData.money }}</p>
                </div>
              
              </el-card>
            </div>
            </el-row>
          </div>
          <div></div>
        </div>
        <img src="@/assets/dashboard.png" class="gva-top-card-right" alt />
      </div>
    </div>

    <!-- <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>数据统计</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="18">
              <echarts-line />
            </el-col>
            <el-col :xs="24" :sm="6">
              <dashboard-table />
            </el-col>
          </el-row>
        </div>
      </div>
    </div> -->
  </div>
</template>

<script setup>
import EchartsLine from "@/view/dashboard/dashboardCharts/echartsLine.vue";
import DashboardTable from "@/view/dashboard/dashboardTable/dashboardTable.vue";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useWeatherInfo } from "@/view/dashboard/weather.js";

import {
  getAgentPayOrderStaticDay,
  getAgentPayOrderStatic,
  getPayOrderStaticDay,
  getPayOrderStatic,
} from "@/api/payOrder";

const staticData = ref({
  amountDay: 0,
  commissionDay: 0,
  countDay: 0,
  amount:0,
  money:0
})

const weatherInfo = useWeatherInfo();

const router = useRouter();

const toTarget = (name) => {
  router.push({ name });
};

const getStatic = async () => {
  const res = await getPayOrderStaticDay();
  if (res.code == 0) {
    staticData.value.amountDay = res.data.moneyAll;
    staticData.value.commissionDay = res.data.money;
    staticData.value.countDay = res.data.count;
  }

  const res1 = await getPayOrderStatic();


  if (res1.code == 0) {
    staticData.value.amount = res1.data.moneyAll;
    staticData.value.money = res1.data.money;
  }
};

getStatic();
</script>
<script>
export default {
  name: "Dashboard",
};
</script>

<style lang="scss" scoped>
@mixin flex-center {
  display: flex;
  align-items: center;
}

.card-container {
  display: flex;
  justify-content: space-between;
}

.box-card {
  margin: 10px;
  width: 300px;
  text-align: center;
}


.item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 15px;
}

.page {
  background: #f0f2f5;
  padding: 0;
  .gva-card-box {
    padding: 12px 16px;
    & + .gva-card-box {
      padding-top: 0px;
    }
  }
  .gva-card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 26px 30px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
  }
  .gva-top-card {
    height: 260px;
    @include flex-center;
    justify-content: space-between;
    color: #777;
    &-left {
      height: 100%;
      display: flex;
      flex-direction: column;
      &-title {
        font-size: 22px;
        color: #343844;
      }
      &-dot {
        font-size: 16px;
        color: #6b7687;
        margin-top: 24px;
      }
      &-rows {
        // margin-top: 15px;
        margin-top: 18px;
        color: #6b7687;
        width: 600px;
        align-items: center;
      }
      &-item {
        + .gva-top-card-left-item {
          margin-top: 24px;
        }
        margin-top: 14px;
      }
    }
    &-right {
      height: 600px;
      width: 600px;
      margin-top: 28px;
    }
  }
  ::v-deep(.el-card__header) {
    padding: 0;
    border-bottom: none;
  }
  .card-header {
    padding-bottom: 20px;
    border-bottom: 1px solid #e8e8e8;
  }
  .quick-entrance-title {
    height: 30px;
    font-size: 22px;
    color: #333;
    width: 100%;
    border-bottom: 1px solid #eee;
  }
  .quick-entrance-items {
    @include flex-center;
    justify-content: center;
    text-align: center;
    color: #333;
    .quick-entrance-item {
      padding: 16px 28px;
      margin-top: -16px;
      margin-bottom: -16px;
      border-radius: 4px;
      transition: all 0.2s;
      &:hover {
        box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
      }
      cursor: pointer;
      height: auto;
      text-align: center;
      // align-items: center;
      &-icon {
        width: 50px;
        height: 50px !important;
        border-radius: 8px;
        @include flex-center;
        justify-content: center;
        margin: 0 auto;
        i {
          font-size: 24px;
        }
      }
      p {
        margin-top: 10px;
      }
    }
  }
  .echart-box {
    padding: 14px;
  }
}
.dashboard-icon {
  font-size: 20px;
  color: rgb(85, 160, 248);
  width: 30px;
  height: 30px;
  margin-right: 10px;
  @include flex-center;
}
.flex-center {
  @include flex-center;
}

//小屏幕不显示右侧，将登录框居中
@media (max-width: 750px) {
  .gva-card {
    padding: 20px 10px !important;
    .gva-top-card {
      height: auto;
      &-left {
        &-title {
          font-size: 20px !important;
        }
        &-rows {
          margin-top: 15px;
          align-items: center;
        }
      }
      &-right {
        display: none;
      }
    }
    .gva-middle-card {
      &-item {
        line-height: 20px;
      }
    }
    .dashboard-icon {
      font-size: 18px;
    }
  }
}
</style>
