# frozen_string_literal: true

module RunApi
  module Gpt4oImage
    module Types
      MODELS = %w[gpt4o-image].freeze
      SIZES = %w[1:1 3:2 2:3].freeze
      FALLBACK_MODELS = %w[GPT_IMAGE_1 FLUX_MAX].freeze
      N_VARIANTS = [ 1, 2, 4 ].freeze

      class Image < RunApi::Core::BaseModel
        optional :url, String
      end

      class GenerationResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :images, [ -> { Image } ]
        optional :progress, String
        optional :error, String
      end

      CompletedGenerationResponse = GenerationResponse
    end
  end
end
