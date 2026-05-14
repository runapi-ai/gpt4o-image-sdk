import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { HttpClient } from '@runapi.ai/core';
import { Generations } from '../../src/resources/generations';
import type { GenerationResponse, TaskCreateResponse } from '../../src/types';

describe('Gpt4oImage Generations', () => {
  const mockHttp: HttpClient = {
    request: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('creates a generation task with direct service params', async () => {
    const mockResponse: TaskCreateResponse = { id: 'task-gen-123', status: 'processing' };
    vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

    const generations = new Generations(mockHttp);
    const result = await generations.create({
      model: 'gpt4o-image',
      prompt: 'A still life',
      size: '1:1',
      files_url: ['https://example.com/input.png'],
      n_variants: 2,
      enable_fallback: true,
      fallback_model: 'FLUX_MAX',
    });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/gpt4o_image/generations', {
      body: {
        model: 'gpt4o-image',
        prompt: 'A still life',
        size: '1:1',
        files_url: ['https://example.com/input.png'],
        n_variants: 2,
        enable_fallback: true,
        fallback_model: 'FLUX_MAX',
      },
    });
    expect(result).toEqual(mockResponse);
  });

  it('gets a generation task', async () => {
    const mockResponse: GenerationResponse = {
      id: 'task-gen-456',
      status: 'completed',
      progress: '1.00',
      images: [{ url: 'https://file.runapi.ai/result.png' }],
    };
    vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

    const generations = new Generations(mockHttp);
    const result = await generations.get('task-gen-456');

    expect(mockHttp.request).toHaveBeenCalledWith('GET', '/api/v1/gpt4o_image/generations/task-gen-456', {});
    expect(result).toEqual(mockResponse);
  });
});
