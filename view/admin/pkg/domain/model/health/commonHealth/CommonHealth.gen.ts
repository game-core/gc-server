import type { CommonHealthEnum } from "~/pkg/domain/model/health/commonHealth/CommonHealthEnum.gen";

export type CommonHealth = {
  healthId: number;
  name: string;
  commonHealthEnum: CommonHealthEnum;
};
