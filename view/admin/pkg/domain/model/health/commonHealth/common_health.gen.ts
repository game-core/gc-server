import type { CommonHealthEnum } from "~/pkg/domain/model/health/commonHealth/common_health_enum.gen";

export type CommonHealth = {
  healthId: number;
  name: string;
  commonHealthEnum: CommonHealthEnum;
};
