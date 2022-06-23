import { IHighlightLineRange } from '@sourcegraph/shared/src/graphql/schema'
import { toAbsoluteBlobURL } from '@sourcegraph/shared/src/util/url'

import { parseBrowserRepoURL } from '../../util/url'

import { Block, BlockInput, FileBlockInput } from '.'

export function serializeBlockInput(block: BlockInput, sourcegraphURL: string): string {
    switch (block.type) {
        case 'md':
        case 'query':
            return block.input
        case 'file':
            return toAbsoluteBlobURL(sourcegraphURL, {
                repoName: block.input.repositoryName,
                revision: block.input.revision,
                filePath: block.input.filePath,
                range: block.input.lineRange
                    ? {
                          start: { line: block.input.lineRange.startLine + 1, character: 0 },
                          end: { line: block.input.lineRange.endLine, character: 0 },
                      }
                    : undefined,
            })
    }
}

function parseFileBlockInput(input: string): FileBlockInput {
    try {
        const { repoName, rawRevision, filePath, position, range } = parseBrowserRepoURL(input)

        const lineRange = range
            ? { startLine: range.start.line - 1, endLine: range.end.line }
            : position
            ? { startLine: position.line - 1, endLine: position.line }
            : null

        return {
            repositoryName: repoName,
            revision: rawRevision ?? '',
            filePath: filePath ?? '',
            lineRange,
        }
    } catch {
        return {
            repositoryName: '',
            revision: '',
            filePath: '',
            lineRange: null,
        }
    }
}

export function deserializeBlockInput(type: Block['type'], input: string): BlockInput {
    switch (type) {
        case 'md':
        case 'query':
            return { type, input }
        case 'file':
            return { type, input: parseFileBlockInput(input) }
    }
}

export function isSingleLineRange(lineRange: IHighlightLineRange | null): boolean {
    return lineRange ? lineRange.startLine + 1 === lineRange.endLine : false
}

export function serializeLineRange(lineRange: IHighlightLineRange | null): string {
    if (!lineRange) {
        return ''
    }

    return isSingleLineRange(lineRange)
        ? `${lineRange.startLine + 1}`
        : `${lineRange.startLine + 1}-${lineRange.endLine}`
}

const LINE_RANGE_REGEX = /^(\d+)(-\d+)?$/

export function parseLineRange(value: string): IHighlightLineRange | null {
    const matches = value.match(LINE_RANGE_REGEX)
    if (matches === null) {
        return null
    }
    const startLine = parseInt(matches[1], 10) - 1
    const endLine = matches[2] ? parseInt(matches[2].slice(1), 10) : startLine + 1
    return { startLine, endLine }
}