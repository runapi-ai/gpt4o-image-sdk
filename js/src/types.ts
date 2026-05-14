import type { AsyncTaskStatus } from '@runapi.ai/core';

export type Gpt4oImageModel = 'gpt4o-image';
export type ImageSize = '1:1' | '3:2' | '2:3';
export type VariantCount = 1 | 2 | 4;
export type FallbackModel = 'GPT_IMAGE_1' | 'FLUX_MAX';

export interface GenerationParams {
  model: Gpt4oImageModel;
  prompt?: string;
  size: ImageSize;
  files_url?: string[];
  mask_url?: string;
  n_variants?: VariantCount;
  callback_url?: string;
  is_enhance?: boolean;
  upload_cn?: boolean;
  enable_fallback?: boolean;
  fallback_model?: FallbackModel;
  file_url?: string;
}

export interface TaskCreateResponse {
  id: string;
  status: string;
}

export interface Image {
  url: string;
}

export interface GenerationResponse {
  id: string;
  status: AsyncTaskStatus;
  images?: Image[];
  progress?: string;
  error?: string;
  [key: string]: unknown;
}
