# AI agent PR review guidelines

Repo-specific conventions for automated PR review (Codex, Cursor, Claude, and any other AI reviewer).

The patterns below often look suspicious in isolation, but they are normal outcomes of Sei's release and upgrade flow. Do not report them unless the diff gives you a separate, concrete reason to believe there is a real bug.

## 1. A release tag that does not exist yet is not a bug

Sei adds release tags in a follow-up step after the dependent code has already merged to `main`. That means it is normal for a PR to introduce a version-gated constant, such as `const FooUpgrade = "v6.7"`, before that version appears in `app/tags`, before a handler exists in `app/upgrades.go`, and before `ctx.ClosestUpgradeName()` or `ctx.LatestUpgrade()` can return it.

Do not report this as "the tag/handler is missing," "the branch is dead," or "the constant does not match anything in the tree." In Sei's workflow, that wiring is often expected to land later with the same version string and the existing `app/tags` naming pattern.

Only treat it as a finding if the diff adds an extra signal, such as:

- the PR explicitly claims to be the release-cut PR and the tag or handler is still missing,
- the new version string does not match Sei's existing tag naming pattern,
- the same feature is tied to inconsistent version strings in different places, or
- the gate clearly points at the wrong upgrade identifier.

## 2. Version gates are not supposed to defend against impossible replay paths

Do not flag a version-gated branch because "a newer binary might process pre-upgrade state" or because "the gate cannot activate yet." That is not how Sei's deployment model works. Binaries move forward version by version: a node syncs with the older binary until the next upgrade height, halts, switches binaries, and continues from there. New code does not replay older not-yet-upgraded state.

The following are not findings by themselves:

- "this branch can never run because the upgrade has not happened yet,"
- "a newer binary could diverge on historical state," or
- "this needs extra compatibility code for older heights."

A gate is only worth reporting if the gate itself is wrong, such as:

- the comparison is inverted,
- the height boundary is off by one,
- the wrong context field is used,
- the semver comparison is wrong,
- one code path is gated and the equivalent path is not, or
- a migration that should happen at the upgrade height does not happen in the handler that claims to own it.

## 3. What a real finding looks like

If you report an upgrade-related issue, make it specific and hard to dispute.

A good finding names the exact constant, gate, handler, or tag; points to the line or lines in the diff that prove the issue; explains why it is a logic bug rather than normal staged release behavior; and suggests the smallest safe fix if there is one.

If the issue disappears once the normal follow-up release tag lands, it is probably not a finding for the PR.

## 4. What to avoid

Do not submit review comments based on:

- assumed future work,
- hypothetical release order changes,
- a missing tag that is expected to be added later,
- a newer binary replaying old state, or
- conclusions of "dead code" that only come from looking at the tree before the release-cut commit lands.

Those are the exact false positives this guide is supposed to prevent.

## 5. Review checklist

Before leaving an upgrade-related comment, ask:

- Is the gate actually wrong, or just not wired yet?
- Does the diff prove a bug, or am I inferring one from normal release staging?
- Is there a concrete mismatch in version string, height, accessor, or comparison?
- Are equivalent code paths handled consistently?
- Would this still be a bug after the usual follow-up release wiring lands?

If the answer is mostly “no,” skip the comment.
